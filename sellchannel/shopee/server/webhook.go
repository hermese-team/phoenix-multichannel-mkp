package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const orderRawTopic = "shopee.order.raw"

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

	// Publish raw event to Kafka asynchronously.
	// The order consumer will fetch full order detail from Shopee API.
	go func() {
		if s.producer == nil {
			return
		}
		raw := OrderRawEvent{
			ShopID:    shopID,
			OrderSN:   orderSN,
			Status:    payload.Data.Status,
			Timestamp: payload.Timestamp,
		}
		msgBytes, _ := json.Marshal(raw)
		if err := s.producer.Publish(context.Background(), orderRawTopic, []byte(orderSN), msgBytes); err != nil {
			fmt.Printf("[webhook] kafka publish raw %s: %v\n", orderSN, err)
		}
	}()
}

func (s *Server) handleProductWebhook(c *gin.Context) {
	// TODO: handle product status change from Shopee
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
