// Package intake implements the W (Webhook Intake) component described in the
// architecture spec section 5.1 Order down-sync steps 6–13.
//
// Both the webhook handler (A→W via HTTP push) and the poll job (A→W via polling)
// route validated raw events through Intake so dedup, safety-net recording, and
// Kafka publish logic are never duplicated.
//
// Step 9 (archive to object storage) is deferred and left as a no-op stub.
package intake

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	redisAdapter  "github.com/okdev/marketplace-sync/internal/infrastructure/redis"
	"github.com/okdev/marketplace-sync/pkg/logger"
)

const (
	// OrderRawTopic is the durable quorum topic all raw order events land on.
	OrderRawTopic = "raw.accepted.shopee.v1.dev"

	// SafetyNetKey is the Redis Set that records every order_sn produced to
	// OrderRawTopic within SafetyNetTTL. The poll fallback uses SMISMEMBER
	// against this set to skip orders already handled by the webhook path.
	SafetyNetKey = "safety-net:shopee:processed"
	SafetyNetTTL = 30 * time.Minute // must match poll cursorOverlap

	// DedupTTL is the TTL for the SetNX dedup key set by Accept.
	// Covers the maximum expected retry window for both webhook retries and
	// overlapping poll windows.
	DedupTTL = 24 * time.Hour
)

// RawEvent is the canonical order event passed from A (channel adapter) to W (intake).
// Code=0 indicates a poll-discovered event (no Shopee push notification).
type RawEvent struct {
	ShopID    int64  `json:"shop_id"`
	OrderSN   string `json:"order_sn"`
	Status    string `json:"status"`
	PushCode  int    `json:"code"`
	Timestamp int64  `json:"timestamp"`
}

// Intake is the W (Webhook Intake) component.
// It is shared by the webhook handler and the poll job.
type Intake struct {
	rdb      *redisAdapter.Client
	producer *kafkaAdapter.Producer
}

func New(rdb *redisAdapter.Client, producer *kafkaAdapter.Producer) *Intake {
	return &Intake{rdb: rdb, producer: producer}
}

// Accept processes a validated RawEvent from A.
//
// pushID is the caller-supplied dedup key:
//   - Webhook path: "shopee:webhook:dedup:{orderSN}:{timestamp}"
//   - Poll path:    "shopee:poll:seen:{orderSN}"
//
// Returns (alreadyProcessed=true, nil) when pushID was already seen — the
// caller should respond 200 OK without further processing (spec step 7).
// Returns (false, err) on Kafka publish failure — the caller should still
// return 200 to Shopee to avoid retry storms; the poll fallback will recover.
func (i *Intake) Accept(ctx context.Context, pushID string, evt RawEvent) (bool, error) {
	// Step 6: Redis fast-path dedup check by push_id.
	if i.rdb != nil && pushID != "" {
		isNew, err := i.rdb.SetNX(ctx, pushID, "1", DedupTTL)
		if err != nil {
			logger.WarnContext(ctx, "intake dedup check failed, proceeding",
				"event", "intake.dedup.error",
				"push_id", pushID,
				"order_sn", evt.OrderSN,
				"error", err,
			)
		} else if !isNew {
			// Step 7: Hit — already processed.
			logger.DebugContext(ctx, "intake dedup hit — already processed",
				"event", "intake.dedup.hit",
				"push_id", pushID,
				"order_sn", evt.OrderSN,
			)
			return true, nil
		}
	}
	// Step 8: Miss — proceed.

	// Step 9: Archive raw payload to object storage — deferred.

	// Step 10: SADD to safety-net set so the poll fallback SMISMEMBER can skip
	// this order and the safety-net scan counts it as matched.
	if i.rdb != nil {
		if err := i.rdb.SAddWithTTL(ctx, SafetyNetKey, SafetyNetTTL, evt.OrderSN); err != nil {
			// Non-fatal — poll fallback may double-process, but idempotency handles it.
			logger.WarnContext(ctx, "intake safety-net sadd failed",
				"event", "intake.safety_net.error",
				"order_sn", evt.OrderSN,
				"error", err,
			)
		}
	}

	// Step 11: Produce to Kafka quorum topic (acks=all).
	msgBytes, _ := json.Marshal(evt)
	if err := i.producer.Publish(ctx, OrderRawTopic, []byte(evt.OrderSN), msgBytes); err != nil {
		return false, fmt.Errorf("intake kafka publish: %w", err)
	}

	// Step 12: Quorum committed — logged by the Kafka producer.

	logger.InfoContext(ctx, "intake accepted",
		"event", "intake.accepted",
		"order_sn", evt.OrderSN,
		"shop_id", evt.ShopID,
		"push_code", evt.PushCode,
	)

	return false, nil
}
