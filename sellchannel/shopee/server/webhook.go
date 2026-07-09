package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/intake"
)

// Shopee push codes dispatched at /webhook.
const (
	pushCodeShopAuthCanceled = 2
	publishTimeout           = 4 * time.Second // Shopee expects response within ~5s
)

// baseWebhookRequest contains the fields common to every Shopee push notification.
// Used by handleWebhook to dispatch to the correct handler before the body is consumed.
type baseWebhookRequest struct {
	Code   int   `json:"code"`
	ShopID int64 `json:"shop_id"`
}

// handleWebhook is the single entry-point for all Shopee push notifications.
// It peeks at the push code and dispatches to the appropriate handler.
func (s *Server) handleWebhook(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}
	// Restore the body so downstream handlers can ShouldBindJSON normally.
	c.Request.Body = io.NopCloser(bytes.NewReader(body))

	var base baseWebhookRequest
	if err := json.Unmarshal(body, &base); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}

	switch base.Code {
	case pushCodeShopAuthCanceled:
		s.handleShopAuthCanceled(c, base.ShopID)
	default:
		// Codes 3+ (order status, product, etc.) go to the order webhook handler.
		s.handleOrderWebhook(c)
	}
}

func (s *Server) handleOrderWebhook(c *gin.Context) {
	var payload OrderWebhookRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		// Still return 200 — Shopee must get 2xx or it will retry indefinitely.
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}

	orderSN := payload.Data.OrderSN
	shopID := payload.ShopID
	reqCtx := c.Request.Context()

	logger.InfoContext(reqCtx, "webhook received",
		"event", "webhook.received",
		"code", payload.Code,
		"shop_id", shopID,
		"order_sn", orderSN,
		"status", payload.Data.Status,
	)

	// Verification ping — Shopee sends an empty payload to confirm the URL.
	if orderSN == "" || shopID == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}

	if s.intake == nil {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}

	// Use a bounded timeout so we always respond before Shopee's deadline (~5s).
	ctx, cancel := context.WithTimeout(reqCtx, publishTimeout)
	defer cancel()

	// Step 3 (A): Skip shops that have not authorized this app.
	if _, _, err := s.tokens.Get(ctx, shopID); err != nil {
		logger.WarnContext(ctx, "webhook shop not authorized, skipping",
			"event", "webhook.shop.unauthorized",
			"shop_id", shopID,
			"order_sn", orderSN,
		)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}

	// Step 4 (A): Build canonical RawEvent.
	evt := intake.RawEvent{
		ShopID:    shopID,
		OrderSN:   orderSN,
		Status:    payload.Data.Status,
		PushCode:  payload.Code,
		Timestamp: payload.Timestamp,
	}

	// Step 5 (A→W): Forward to webhook intake.
	// pushID includes timestamp so each Shopee retry attempt has a unique key.
	pushID := fmt.Sprintf("shopee:webhook:dedup:%s:%d", orderSN, payload.Timestamp)
	alreadyProcessed, err := s.intake.Accept(ctx, pushID, evt)
	if err != nil {
		logger.ErrorContext(ctx, "intake failed",
			"event", "webhook.intake.error",
			"order_sn", orderSN,
			"shop_id", shopID,
			"error", err,
		)
		// Return 200 to avoid Shopee flooding during a Kafka outage.
		// The poll fallback scheduler will catch missed orders.
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}
	if alreadyProcessed {
		// Step 7: dedup hit — return 200 so Shopee stops retrying.
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}

	// Steps 13-14: W accepted → A returns 202 to Shopee.
	c.JSON(http.StatusAccepted, gin.H{"status": "accepted"})
}

func (s *Server) handleProductWebhook(c *gin.Context) {
	// TODO: handle product status change from Shopee
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
