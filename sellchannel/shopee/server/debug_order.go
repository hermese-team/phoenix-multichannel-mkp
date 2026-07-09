package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/intake"
)

type debugFetchOrderRequest struct {
	ShopID  int64  `form:"shop_id" binding:"required"`
	OrderSN string `form:"order_sn" binding:"required"`
}

// GET /debug/order?shop_id=225997847&order_sn=260629CC7NGG5Y
// Simulates a webhook push: fetches the order from Shopee API to verify the token,
// then publishes a raw event to order.raw.accepted.v1 for the consumer to enrich.
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

	// 2. publish raw event to order.raw.accepted.v1 — consumer will fetch full detail
	if s.intake != nil {
		evt := intake.RawEvent{
			ShopID:    req.ShopID,
			OrderSN:   req.OrderSN,
			Status:    "DEBUG",
			Timestamp: time.Now().Unix(),
		}
		pushID := fmt.Sprintf("shopee:debug:%s:%d", req.OrderSN, evt.Timestamp)
		if _, err := s.intake.Accept(c.Request.Context(), pushID, evt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "kafka publish: " + err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"shop_id":         req.ShopID,
		"order_sn":        req.OrderSN,
		"kafka_topic":     intake.OrderRawTopic,
		"kafka_published": s.intake != nil,
		"message":         "raw event published — consumer will fetch order detail from Shopee API",
	})
}
