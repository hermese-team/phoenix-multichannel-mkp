package server

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	orderUC "github.com/okdev/marketplace-sync/internal/usecase/order"
	"github.com/okdev/marketplace-sync/pkg/apperror"
	"github.com/okdev/marketplace-sync/pkg/httpserver"
	applog "github.com/okdev/marketplace-sync/pkg/logger"
)

// idemTTL is how long a processed message key is remembered in Redis.
const idemTTL = 24 * time.Hour

func (s *Server) handleOrderWebhook(c *gin.Context) {
	// อ่าน raw body ครั้งเดียว แล้ว log ดูทั้งก้อนที่ Lazada ยิงมา
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		respondError(c, apperror.NewBadRequest(err))
		return
	}
	// NOTE: logs the full raw body at Info. Fine today (the payload has no
	// PII/secret), but revisit if the webhook schema grows to include personal data.
	applog.Info("lazada webhook received", "body", string(body))

	// parse จาก body ที่อ่านไว้ (ไม่ใช้ ShouldBindJSON เพราะ body ถูกอ่านไปแล้ว)
	var payload OrderWebhookRequest
	if err := json.Unmarshal(body, &payload); err != nil {
		respondError(c, apperror.NewBadRequest(err))
		return
	}

	// Idempotency in Redis: dedup the same status transition (order + update
	// time). Duplicate pushes are acked with 200 but not reprocessed. On a Redis
	// error we fail open (process anyway) so messages are not lost.
	idemKey := fmt.Sprintf("%s:%s:%d", s.idemPrefix, payload.Data.TradeOrderID, payload.Data.StatusUpdateTime)
	if isNew, err := s.redis.SetNX(c.Request.Context(), idemKey, "1", idemTTL); err != nil {
		applog.Error("idempotency check", "key", idemKey, "error", err)
	} else if !isNew {
		applog.Info("duplicate webhook skipped", "order", payload.Data.TradeOrderID)
		httpserver.NewSuccessResponse(c, gin.H{"status": "ok"})
		return
	}

	// Publish the raw message to Kafka (key = order id) for downstream consumers.
	// Best-effort: a publish failure is logged but does not fail the webhook.
	if err := s.producer.Publish(c.Request.Context(), s.topic, []byte(payload.Data.TradeOrderID), body); err != nil {
		applog.Error("publish webhook to kafka", "topic", s.topic, "order", payload.Data.TradeOrderID, "error", err)
	}

	// Ack Lazada fast with 200 and process best-effort — a webhook must not
	// fail verification/delivery just because downstream persistence errored.
	// seller_id arrives as a string; a non-numeric one (the verify test uses
	// "9999") still parses, others fall back to 0.
	shopID, _ := strconv.ParseInt(payload.SellerID, 10, 64)

	// NOTE(mock): the webhook both publishes to Kafka (the consumer does the full
	// async detail sync) and writes the order synchronously below. This
	// double-handling is a mock convenience — revisit once the async pipeline is
	// the single source of truth for persistence.
	if err := s.processOrder.Execute(c.Request.Context(), orderUC.WebhookPayload{
		MarketplaceOrderID: payload.Data.TradeOrderID,
		SellChannelType:    "lazada",
		ShopID:             shopID,
		Status:             payload.Data.OrderStatus,
	}); err != nil {
		applog.Error("process lazada order webhook", "order", payload.Data.TradeOrderID, "error", err)
	}

	httpserver.NewSuccessResponse(c, gin.H{"status": "ok"})
}
