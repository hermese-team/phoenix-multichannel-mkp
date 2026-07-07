package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	pgAdapter     "github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	redisAdapter  "github.com/okdev/marketplace-sync/internal/infrastructure/redis"
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
		log.Printf("[order-poll] find shop ids: %v", err)
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
		log.Printf("[order-poll] no token for shop %d: %v", shopID, err)
		return
	}

	orders, err := j.shopeeClient.GetOrderList(ctx, shopID, accessToken, timeFrom, timeTo)
	if err != nil {
		log.Printf("[order-poll] get order list shop %d: %v", shopID, err)
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
			log.Printf("[order-poll] publish %s: %v", o.OrderSN, err)
			continue
		}
		published++
	}

	if published > 0 {
		log.Printf("[order-poll] shop %d: published %d/%d orders", shopID, published, len(orders))
	}
}
