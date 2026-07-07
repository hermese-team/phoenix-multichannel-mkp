package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type debugFetchOrderRequest struct {
	ShopID  int64  `form:"shop_id" binding:"required"`
	OrderSN string `form:"order_sn" binding:"required"`
}

// GET /debug/order?shop_id=225997847&order_sn=260629CC7NGG5Y
// Simulates a webhook push: fetches the order from Shopee API to verify the token,
// then publishes a raw event to shopee.order.raw for the consumer to enrich.
func (s *Server) debugFetchOrder(c *gin.Context) {
	var req debugFetchOrderRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. verify token exists
	accessToken, _, err := s.tokens.Get(c.Request.Context(), req.ShopID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("no token for shop %d: %v", req.ShopID, err)})
		return
	}
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("shop %d not authorized", req.ShopID)})
		return
	}

	// 2. publish raw event to shopee.order.raw — consumer will fetch full detail
	raw := OrderRawEvent{
		ShopID:    req.ShopID,
		OrderSN:   req.OrderSN,
		Status:    "DEBUG",
		Timestamp: time.Now().Unix(),
	}
	if s.producer != nil {
		payload, _ := json.Marshal(raw)
		if err := s.producer.Publish(c.Request.Context(), orderRawTopic, []byte(req.OrderSN), payload); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "kafka publish: " + err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"shop_id":         req.ShopID,
		"order_sn":        req.OrderSN,
		"kafka_topic":     orderRawTopic,
		"kafka_published": s.producer != nil,
		"message":         "raw event published — consumer will fetch order detail from Shopee API",
	})
}
