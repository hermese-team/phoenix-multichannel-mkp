package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) handleOrderWebhook(c *gin.Context) {
	var payload OrderWebhookRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		// Still return 200 — Shopee must get 2xx or it will retry indefinitely.
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}

	// Acknowledge immediately — Shopee requires 2xx within a short timeout.
	c.JSON(http.StatusOK, gin.H{"status": "ok"})

	orderSN := payload.Data.OrderSN
	shopID := payload.ShopID
	fmt.Printf("[webhook] received code=%d shop_id=%d order_sn=%s status=%s\n",
		payload.Code, shopID, orderSN, payload.Data.Status)
	if orderSN == "" || shopID == 0 {
		return // verification ping — nothing to process
	}

	// Process asynchronously so we never block Shopee's push timeout.
	go func() {
		ctx := context.Background()

		accessToken, _, err := s.tokens.Get(ctx, shopID)
		if err != nil {
			fmt.Printf("[webhook] no token for shop %d: %v\n", shopID, err)
			return
		}
		if accessToken == "" {
			fmt.Printf("[webhook] shop %d not authorized\n", shopID)
			return
		}

		orders, err := s.shopeeClient.GetOrderDetail(ctx, shopID, accessToken, []string{orderSN})
		if err != nil {
			fmt.Printf("[webhook] fetch order %s: %v\n", orderSN, err)
			return
		}
		if len(orders) == 0 {
			fmt.Printf("[webhook] order %s not found in Shopee\n", orderSN)
			return
		}

		order := orders[0]
		order.ShopID = shopID // Shopee API omits shop_id in order detail response

		if s.producer != nil {
			msgBytes, _ := json.Marshal(order)
			if err := s.producer.Publish(ctx, orderWebhookTopic, []byte(orderSN), msgBytes); err != nil {
				fmt.Printf("[webhook] kafka publish %s: %v\n", orderSN, err)
			}
		}
	}()
}

func (s *Server) handleProductWebhook(c *gin.Context) {
	// TODO: handle product status change from Shopee
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
