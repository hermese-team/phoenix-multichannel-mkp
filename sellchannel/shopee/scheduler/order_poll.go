package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	pgAdapter    "github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	redisAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/redis"
	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/client"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/tokenstore"
)

const (
	orderRawTopic = "shopee.order.raw"

	// safetyNetKey must match the key written by the webhook handler.
	safetyNetKey = "safety-net:shopee:processed"

	// cursorOverlap is the minimum look-back window for each poll run.
	// Must be >= safetyNetTTL (30 min) so the SMISMEMBER cross-reference
	// covers every order that the webhook could have marked.
	cursorOverlap = 30 * time.Minute

	// pollDedupTTL is the SetNX TTL for poll-level idempotency.
	// Slightly longer than cursorOverlap so a seen key doesn't expire
	// before the next overlapping window.
	pollDedupTTL = 35 * time.Minute

	// leaseTTL is the distributed lock TTL.
	// Set to slightly less than the poll interval (e.g. 4 min for @every 5m)
	// so the lock auto-expires if a pod crashes mid-poll.
	leaseTTL = 4 * time.Minute
	leaseKey = "shopee:poll:lease"
)

// orderRawEvent mirrors server.OrderRawEvent — defined here to avoid circular imports.
type orderRawEvent struct {
	ShopID    int64  `json:"shop_id"`
	OrderSN   string `json:"order_sn"`
	Status    string `json:"status"`
	Code      int    `json:"code"` // 0 = poll-discovered (no push notification)
	Timestamp int64  `json:"timestamp"`
}

// OrderPollJob polls Shopee order list for all authorized shops and publishes
// any missed orders to shopee.order.raw so the consumer can enrich them.
//
// Safety layers (innermost to outermost):
//  1. SMISMEMBER on safety-net:shopee:processed — skip webhook-handled orders
//  2. SetNX shopee:poll:seen:{sn}              — poll-level dedup
type OrderPollJob struct {
	shopeeClient *client.Client
	tokens       *tokenstore.Store
	tokenRepo    *pgAdapter.ShopeeTokenRepository
	producer     *kafkaAdapter.Producer
	rdb          *redisAdapter.Client
	window       time.Duration
}

func NewOrderPollJob(
	shopeeClient *client.Client,
	tokens *tokenstore.Store,
	tokenRepo *pgAdapter.ShopeeTokenRepository,
	producer *kafkaAdapter.Producer,
	rdb *redisAdapter.Client,
	window time.Duration,
) *OrderPollJob {
	return &OrderPollJob{
		shopeeClient: shopeeClient,
		tokens:       tokens,
		tokenRepo:    tokenRepo,
		producer:     producer,
		rdb:          rdb,
		window:       window,
	}
}

// Run is called by the cron scheduler on each tick.
func (j *OrderPollJob) Run() {
	ctx := context.Background()

	// Distributed lease — only one pod polls at a time.
	if j.rdb != nil {
		acquired, err := j.rdb.SetNX(ctx, leaseKey, "1", leaseTTL)
		if err != nil {
			logger.WarnContext(ctx, "poll lease check failed, proceeding anyway",
				"event", "order_poll.lease.error",
				"error", err,
			)
		} else if !acquired {
			logger.DebugContext(ctx, "poll skipped: lease held by another instance",
				"event", "order_poll.lease.skipped",
			)
			return
		}
		// Release lock when done so the next scheduled tick can acquire it immediately.
		defer j.rdb.Delete(ctx, leaseKey)
	}

	shopIDs, err := j.tokenRepo.FindAllShopIDs(ctx)
	if err != nil {
		logger.ErrorContext(ctx, "find shop ids failed",
			"event", "order_poll.shop_ids.error",
			"error", err,
		)
		return
	}
	if len(shopIDs) == 0 {
		return
	}

	// Overlap-safe cursor: look back at least cursorOverlap regardless of the
	// configured window, so we never miss orders that arrived near the boundary
	// of the previous poll window.
	lookback := j.window
	if lookback < cursorOverlap {
		lookback = cursorOverlap
	}
	now := time.Now()
	timeFrom := now.Add(-lookback).Unix()
	timeTo := now.Unix()

	logger.InfoContext(ctx, "poll run started",
		"event", "order_poll.run.started",
		"shops", len(shopIDs),
		"lookback_min", int(lookback.Minutes()),
	)

	for _, shopID := range shopIDs {
		j.pollShop(ctx, shopID, timeFrom, timeTo)
	}
}

func (j *OrderPollJob) pollShop(ctx context.Context, shopID, timeFrom, timeTo int64) {
	accessToken, _, err := j.tokens.Get(ctx, shopID)
	if err != nil || accessToken == "" {
		logger.WarnContext(ctx, "no token for shop",
			"event", "order_poll.token.missing",
			"shop_id", shopID,
			"error", err,
		)
		return
	}

	orders, err := j.shopeeClient.GetOrderList(ctx, shopID, accessToken, timeFrom, timeTo)
	if err != nil {
		logger.ErrorContext(ctx, "get order list failed",
			"event", "order_poll.get_order_list.error",
			"shop_id", shopID,
			"error", err,
		)
		return
	}
	if len(orders) == 0 {
		return
	}

	// ── Layer 1: SMISMEMBER cross-reference with webhook safety-net ──────────
	// Any order already in the safety-net was handled by the webhook path within
	// the last 30 minutes — skip it to avoid double-processing.
	if j.rdb != nil {
		sns := make([]string, len(orders))
		for i, o := range orders {
			sns[i] = o.OrderSN
		}
		inSafetyNet, err := j.rdb.SMIsMember(ctx, safetyNetKey, sns...)
		if err != nil {
			logger.WarnContext(ctx, "safety-net SMISMEMBER failed, falling back to SetNX only",
				"event", "order_poll.safety_net.error",
				"shop_id", shopID,
				"error", err,
			)
		} else {
			filtered := orders[:0]
			skipped := 0
			for i, o := range orders {
				if inSafetyNet[i] {
					skipped++
				} else {
					filtered = append(filtered, o)
				}
			}
			if skipped > 0 {
				logger.DebugContext(ctx, "poll: safety-net skipped",
					"event", "order_poll.safety_net.skip",
					"shop_id", shopID,
					"skipped", skipped,
				)
			}
			orders = filtered
		}
	}

	if len(orders) == 0 {
		return
	}

	// ── Layer 2: SetNX poll-level dedup ──────────────────────────────────────
	// Prevents the same order from being published multiple times within a
	// single overlapping poll window (e.g. two overlapping 30-min windows).
	published := 0
	for _, o := range orders {
		if j.rdb != nil {
			key := fmt.Sprintf("shopee:poll:seen:%s", o.OrderSN)
			isNew, err := j.rdb.SetNX(ctx, key, "1", pollDedupTTL)
			if err == nil && !isNew {
				continue
			}
		}

		raw := orderRawEvent{
			ShopID:    shopID,
			OrderSN:   o.OrderSN,
			Status:    o.OrderStatus,
			Code:      0, // poll-discovered; classifier handles code=0 via status
			Timestamp: time.Now().Unix(),
		}
		msgBytes, _ := json.Marshal(raw)
		if err := j.producer.Publish(ctx, orderRawTopic, []byte(o.OrderSN), msgBytes); err != nil {
			logger.ErrorContext(ctx, "publish poll order failed",
				"event", "order_poll.kafka.error",
				"order_sn", o.OrderSN,
				"shop_id", shopID,
				"error", err,
			)
			continue
		}
		published++
	}

	if published > 0 {
		logger.InfoContext(ctx, "poll published orders",
			"event", "order_poll.published",
			"shop_id", shopID,
			"published", published,
			"total", len(orders),
		)
	}
}
