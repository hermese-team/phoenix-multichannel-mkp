package main

// webhookHandler receives Shopee push notification events and publishes them
// to the Kafka raw order topic (raw.order.accepted.v1).
//
// Architecture contract (from architecture01-pmc.md §4.1 Channel ingress plane):
//   - Verify Shopee HMAC-SHA256 signature before any processing.
//   - Enforce a payload size limit (1 MiB) to protect downstream consumers.
//   - Produce the raw payload to Kafka with acks=all (ProduceSync).
//   - Return 200 only after Kafka confirms quorum — never before.
//   - This handler must NOT call PostgreSQL, Redis, or external APIs.
//
// Shopee webhook signature (Open Platform v2):
//   Authorization header = "{partner_id}.{signature}"
//   signature = lower-hex HMAC-SHA256(partner_key, partner_id + push_url + raw_body)

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"git.amaze-x.com/phoenix/channel-adapter-shopee/config"
	"git.amaze-x.com/phoenix/channel-adapter-shopee/internal/infrastructure"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const maxWebhookBodyBytes = 1 << 20 // 1 MiB

// webhookEvent is a minimal envelope used only for Kafka key extraction.
// The full raw body is forwarded to Kafka without modification.
type webhookEvent struct {
	Code   int    `json:"code"`
	ShopID int64  `json:"shop_id"`
	Data   struct {
		OrderSN string `json:"ordersn"`
	} `json:"data"`
}

func webhookHandler(cfg config.Config, producer *infrastructure.KafkaProducer, log *zap.Logger) gin.HandlerFunc {
	partnerID := cfg.Shopee.PartnerID
	partnerKey := cfg.Shopee.PartnerKey
	topic := cfg.Kafka.TopicRawOrderAccepted

	return func(c *gin.Context) {
		// ── 1. Size guard ────────────────────────────────────────────────────
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxWebhookBodyBytes)
		rawBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Warn("webhook body read failed", zap.Error(err))
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "payload too large"})
			return
		}

		// ── 2. Signature verification ─────────────────────────────────────────
		// Authorization: {partner_id}.{hmac_hex}
		authHeader := c.GetHeader("Authorization")
		if !verifyShopeeSignature(authHeader, partnerID, partnerKey, c.Request.RequestURI, rawBody) {
			log.Warn("webhook signature mismatch",
				zap.String("remote_addr", c.ClientIP()),
				zap.String("auth_header", authHeader),
			)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid signature"})
			return
		}

		// ── 3. Produce to Kafka (acks=all, blocks until quorum) ──────────────
		// Partition key: shop_id + ordersn for per-order ordering guarantee.
		// We extract just enough to build the key; the full raw payload is forwarded.
		shopID := fmt.Sprintf("%d", cfg.Shopee.ShopID)
		key := shopID // fallback; overridden below if we can parse the order SN

		var evt webhookEvent
		if jsonErr := json.Unmarshal(rawBody, &evt); jsonErr == nil && evt.Data.OrderSN != "" {
			key = fmt.Sprintf("%d-%s", evt.ShopID, evt.Data.OrderSN)
		}

		if err := producer.Produce(c.Request.Context(), topic, key, rawBody); err != nil {
			log.Error("kafka produce failed", zap.Error(err), zap.String("topic", topic))
			// Return 500 so Shopee retries — we must not acknowledge events we failed to store.
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		log.Info("shopee webhook received and committed",
			zap.String("key", key),
			zap.String("topic", topic),
			zap.Int("body_bytes", len(rawBody)),
		)

		// Return 200 after Kafka quorum — Shopee will not retry.
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

// verifyShopeeSignature validates the Shopee webhook Authorization header.
//
// Expected header format: "{partner_id}.{hmac_sha256_hex}"
// Expected HMAC base string: "{partner_id}{request_uri}{raw_body}"
func verifyShopeeSignature(authHeader string, partnerID int64, partnerKey, requestURI string, body []byte) bool {
	if authHeader == "" {
		return false
	}

	// Split "{partner_id}.{sig}"
	dot := strings.Index(authHeader, ".")
	if dot < 0 {
		return false
	}
	receivedSig := authHeader[dot+1:]

	// Recompute expected signature
	base := fmt.Sprintf("%d%s", partnerID, requestURI) + string(body)
	mac := hmac.New(sha256.New, []byte(partnerKey))
	mac.Write([]byte(base))
	expectedSig := hex.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(receivedSig), []byte(expectedSig))
}
