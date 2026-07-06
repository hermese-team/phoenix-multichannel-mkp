package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const orderWebhookTopic = "shopee.order.webhook"

type debugFetchOrderRequest struct {
	ShopID  int64  `form:"shop_id" binding:"required"`
	OrderSN string `form:"order_sn" binding:"required"`
}

// GET /debug/order?shop_id=225997847&order_sn=260629CC7NGG5Y
// Fetches the order from Shopee using the stored token and publishes to Kafka.
func (s *Server) debugFetchOrder(c *gin.Context) {
	var req debugFetchOrderRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. get stored access token
	accessToken, _, err := s.tokens.Get(c.Request.Context(), req.ShopID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("no token for shop %d: %v", req.ShopID, err)})
		return
	}

	// 2. call Shopee GetOrderDetail
	orders, err := s.shopeeClient.GetOrderDetail(c.Request.Context(), req.ShopID, accessToken, []string{req.OrderSN})
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	if len(orders) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	order := orders[0]
	order.ShopID = req.ShopID // Shopee API omits shop_id in order detail response

	// 3. publish to Kafka
	if s.producer != nil {
		payload, _ := json.Marshal(order)
		if err := s.producer.Publish(c.Request.Context(), orderWebhookTopic, []byte(req.OrderSN), payload); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "kafka publish: " + err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"order":           order,
		"kafka_topic":     orderWebhookTopic,
		"kafka_published": s.producer != nil,
	})
}
