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
	orderRawTopic = "shopee.order.raw"
	dedupKeyTTL   = 24 * time.Hour
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
	reqCtx := c.Request.Context()

	logger.InfoContext(reqCtx, "webhook received",
		"event", "webhook.received",
		"code", payload.Code,
		"shop_id", shopID,
		"order_sn", orderSN,
		"status", payload.Data.Status,
	)

	if orderSN == "" || shopID == 0 {
		return // verification ping — nothing to process
	}

	// Publish raw event to Kafka asynchronously.
	// The order consumer will fetch full order detail from Shopee API.
	go func() {
		ctx := context.Background()

		if s.producer == nil {
			return
		}

		// Idempotency: deduplicate retried webhooks.
		// Shopee retries with the same timestamp if it doesn't receive 2xx in time.
		// SetNX returns true = first time seen, false = duplicate.
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
			return
		}

		raw := OrderRawEvent{
			ShopID:    shopID,
			OrderSN:   orderSN,
			Status:    payload.Data.Status,
			Timestamp: payload.Timestamp,
		}
		msgBytes, _ := json.Marshal(raw)
		if err := s.producer.Publish(ctx, orderRawTopic, []byte(orderSN), msgBytes); err != nil {
			logger.ErrorContext(ctx, "webhook kafka publish failed",
				"event", "webhook.kafka.error",
				"order_sn", orderSN,
				"shop_id", shopID,
				"error", err,
			)
		} else {
			logger.InfoContext(ctx, "webhook published to kafka",
				"event", "webhook.kafka.published",
				"order_sn", orderSN,
				"shop_id", shopID,
				"topic", orderRawTopic,
			)
		}
	}()
}

func (s *Server) handleProductWebhook(c *gin.Context) {
	// TODO: handle product status change from Shopee
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
