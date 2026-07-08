package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	pgAdapter     "github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	redisAdapter  "github.com/okdev/marketplace-sync/internal/infrastructure/redis"
	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/client"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/tokenstore"
)

const (
	orderRawTopic  = "shopee.order.raw"
	pollDedupTTL   = 10 * time.Minute // skip order if already seen within this window
)

// orderRawEvent mirrors server.OrderRawEvent — defined here to avoid circular imports.
type orderRawEvent struct {
	ShopID    int64  `json:"shop_id"`
	OrderSN   string `json:"order_sn"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

// OrderPollJob polls Shopee order list for all authorized shops and publishes
// any missed orders to shopee.order.raw so the consumer can enrich them.
type OrderPollJob struct {
	shopeeClient *client.Client
	tokens       *tokenstore.Store
	tokenRepo    *pgAdapter.ShopeeTokenRepository
	producer     *kafkaAdapter.Producer
	rdb          *redisAdapter.Client
	window       time.Duration // how far back to look (should be > poll interval)
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

	now := time.Now()
	timeFrom := now.Add(-j.window).Unix()
	timeTo := now.Unix()

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

	published := 0
	for _, o := range orders {
		// Dedup: skip if already seen in this polling window.
		if j.rdb != nil {
			key := fmt.Sprintf("shopee:poll:seen:%s", o.OrderSN)
			isNew, err := j.rdb.SetNX(ctx, key, "1", pollDedupTTL)
			if err == nil && !isNew {
				continue // already published by webhook or previous poll
			}
		}

		raw := orderRawEvent{
			ShopID:    shopID,
			OrderSN:   o.OrderSN,
			Status:    o.OrderStatus,
			Timestamp: time.Now().Unix(),
		}
		msgBytes, _ := json.Marshal(raw)
		if err := j.producer.Publish(ctx, orderRawTopic, []byte(o.OrderSN), msgBytes); err != nil {
			logger.ErrorContext(ctx, "publish order failed",
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
