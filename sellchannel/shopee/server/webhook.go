package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/okdev/marketplace-sync/pkg/logger"
)

const (
	orderRawTopic  = "shopee.order.raw"
	dedupKeyTTL    = 24 * time.Hour
	safetyNetKey   = "safety-net:shopee:processed"
	safetyNetTTL   = 30 * time.Minute // matches poll fallback scan window
	publishTimeout = 4 * time.Second  // Shopee expects response within ~5s
)

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

	if s.producer == nil {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}

	// Use a bounded timeout so we always respond before Shopee's deadline (~5s).
	ctx, cancel := context.WithTimeout(reqCtx, publishTimeout)
	defer cancel()

	// Idempotency: deduplicate retried webhooks.
	// Shopee retries with the same timestamp if it doesn't receive 2xx in time.
	if s.rdb != nil {
		key := fmt.Sprintf("shopee:webhook:dedup:%s:%d", orderSN, payload.Timestamp)
		isNew, err := s.rdb.SetNX(ctx, key, "1", dedupKeyTTL)
		if err != nil {
			logger.WarnContext(ctx, "webhook dedup check failed, proceeding",
				"event", "webhook.dedup.error",
				"order_sn", orderSN,
				"error", err,
			)
		} else if !isNew {
			logger.DebugContext(ctx, "webhook duplicate skipped",
				"event", "webhook.dedup.skip",
				"order_sn", orderSN,
				"timestamp", payload.Timestamp,
			)
			// Return 200 so Shopee stops retrying this exact push.
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
			return
		}
	}

	// Skip shops that have not authorized this app — no token, nothing to do.
	if _, _, err := s.tokens.Get(ctx, shopID); err != nil {
		logger.WarnContext(ctx, "webhook shop not authorized, skipping",
			"event", "webhook.shop.unauthorized",
			"shop_id", shopID,
			"order_sn", orderSN,
		)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}

	raw := OrderRawEvent{
		ShopID:    shopID,
		OrderSN:   orderSN,
		Status:    payload.Data.Status,
		Code:      payload.Code,
		Timestamp: payload.Timestamp,
	}
	msgBytes, _ := json.Marshal(raw)

	// Synchronous publish — response is sent only after Kafka quorum confirms.
	if err := s.producer.Publish(ctx, orderRawTopic, []byte(orderSN), msgBytes); err != nil {
		logger.ErrorContext(ctx, "webhook kafka publish failed",
			"event", "webhook.kafka.error",
			"order_sn", orderSN,
			"shop_id", shopID,
			"error", err,
		)
		// Return 200 to avoid Shopee flooding during a Kafka outage.
		// The poll fallback scheduler will catch missed orders.
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}

	// Safety-net: record order_sn in a Redis Set so the poll fallback can
	// skip orders already processed here (cross-referenced via SMISMEMBER).
	if s.rdb != nil {
		if err := s.rdb.SAddWithTTL(ctx, safetyNetKey, safetyNetTTL, orderSN); err != nil {
			// Non-fatal — poll fallback will still work, just may double-process.
			logger.WarnContext(ctx, "safety-net sadd failed",
				"event", "webhook.safety_net.error",
				"order_sn", orderSN,
				"error", err,
			)
		}
	}

	logger.InfoContext(ctx, "webhook published to kafka",
		"event", "webhook.kafka.published",
		"order_sn", orderSN,
		"shop_id", shopID,
		"topic", orderRawTopic,
	)

	// 202 Accepted — Kafka quorum confirmed, business processing happens async.
	c.JSON(http.StatusAccepted, gin.H{"status": "accepted"})
}

func (s *Server) handleProductWebhook(c *gin.Context) {
	// TODO: handle product status change from Shopee
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
