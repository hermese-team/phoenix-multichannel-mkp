package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	orderUC "github.com/okdev/marketplace-sync/internal/usecase/order"
)

func (s *Server) handleOrderWebhook(c *gin.Context) {
	var payload OrderWebhookRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.processOrder.Execute(c.Request.Context(), orderUC.WebhookPayload{
		MarketplaceOrderID: payload.OrderSN,
		SellChannelType:    "shopee",
		ShopID:             payload.ShopID,
		Status:             payload.Status,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (s *Server) handleProductWebhook(c *gin.Context) {
	// TODO: handle product status change from Shopee
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
