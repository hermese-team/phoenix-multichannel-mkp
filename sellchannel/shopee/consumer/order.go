package consumer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"

	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	pgAdapter     "github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/client"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/tokenstore"
)

const (
	orderConsumerGroup = "shopee-order-consumer"
	orderRawTopic      = "shopee.order.raw"
	orderDetailTopic   = "shopee.order.detail"
)

// orderRawEvent mirrors server.OrderRawEvent — defined here to avoid circular imports.
type orderRawEvent struct {
	ShopID    int64  `json:"shop_id"`
	OrderSN   string `json:"order_sn"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

// OrderConsumer reads raw webhook events from shopee.order.raw,
// fetches full order detail from Shopee API, and publishes to shopee.order.detail.
// When an order is READY_TO_SHIP it enqueues a fulfillment row for the scheduler.
type OrderConsumer struct {
	cfg             kafkaAdapter.Config
	shopeeClient    *client.Client
	tokens          *tokenstore.Store
	producer        *kafkaAdapter.Producer
	fulfillmentRepo *pgAdapter.ShopeeFulfillmentRepository // may be nil
}

func NewOrderConsumer(cfg kafkaAdapter.Config, shopeeClient *client.Client, tokens *tokenstore.Store, producer *kafkaAdapter.Producer, fulfillmentRepo *pgAdapter.ShopeeFulfillmentRepository) *OrderConsumer {
	return &OrderConsumer{
		cfg:             cfg,
		shopeeClient:    shopeeClient,
		tokens:          tokens,
		producer:        producer,
		fulfillmentRepo: fulfillmentRepo,
	}
}

func (c *OrderConsumer) Start(ctx context.Context) error {
	kafkaClient, err := kgo.NewClient(
		kgo.SeedBrokers(c.cfg.Brokers...),
		kgo.ConsumerGroup(orderConsumerGroup),
		kgo.ConsumeTopics(orderRawTopic),
	)
	if err != nil {
		return err
	}
	defer kafkaClient.Close()

	logger.InfoContext(ctx, "order consumer started",
		"event", "consumer.started",
		"topic", orderRawTopic,
		"group", orderConsumerGroup,
	)
	for {
		if ctx.Err() != nil {
			return nil
		}
		fetches := kafkaClient.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			if ctx.Err() != nil {
				return nil
			}
			for _, e := range errs {
				logger.ErrorContext(ctx, "kafka fetch error",
					"event", "consumer.fetch.error",
					"error", e.Err,
				)
			}
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(time.Second):
			}
			continue
		}
		fetches.EachRecord(func(rec *kgo.Record) {
			var event orderRawEvent
			if err := json.Unmarshal(rec.Value, &event); err != nil {
				logger.ErrorContext(ctx, "unmarshal event failed",
					"event", "consumer.unmarshal.error",
					"error", err,
				)
				return
			}
			c.process(ctx, event)
		})
	}
}

func (c *OrderConsumer) process(ctx context.Context, event orderRawEvent) {
	accessToken, _, err := c.tokens.Get(ctx, event.ShopID)
	if err != nil {
		logger.ErrorContext(ctx, "get token failed",
			"event", "consumer.token.error",
			"shop_id", event.ShopID,
			"error", err,
		)
		return
	}
	if accessToken == "" {
		logger.WarnContext(ctx, "shop not authorized",
			"event", "consumer.shop.unauthorized",
			"shop_id", event.ShopID,
		)
		return
	}

	orders, err := c.shopeeClient.GetOrderDetail(ctx, event.ShopID, accessToken, []string{event.OrderSN})
	if err != nil {
		logger.ErrorContext(ctx, "get order detail failed",
			"event", "consumer.get_order_detail.error",
			"order_sn", event.OrderSN,
			"shop_id", event.ShopID,
			"error", err,
		)
		return
	}
	if len(orders) == 0 {
		logger.WarnContext(ctx, "order not found in shopee",
			"event", "consumer.order.not_found",
			"order_sn", event.OrderSN,
			"shop_id", event.ShopID,
		)
		return
	}

	order := orders[0]
	order.ShopID = event.ShopID // Shopee API omits shop_id in response

	msgBytes, _ := json.Marshal(order)
	if err := c.producer.Publish(ctx, orderDetailTopic, []byte(event.OrderSN), msgBytes); err != nil {
		logger.ErrorContext(ctx, "kafka publish order detail failed",
			"event", "consumer.kafka.error",
			"order_sn", event.OrderSN,
			"shop_id", event.ShopID,
			"error", err,
		)
		return
	}

	logger.InfoContext(ctx, "order detail published",
		"event", "consumer.order.published",
		"order_sn", event.OrderSN,
		"shop_id", event.ShopID,
		"order_status", order.OrderStatus,
		"topic", orderDetailTopic,
	)

	// Enqueue fulfillment when order is ready to ship.
	if order.OrderStatus == "READY_TO_SHIP" && c.fulfillmentRepo != nil {
		if err := c.fulfillmentRepo.Enqueue(ctx, event.ShopID, event.OrderSN,
			pgAdapter.ShopeeDeliveryShopeeLogistics, ""); err != nil {
			logger.ErrorContext(ctx, "enqueue fulfillment failed",
				"event", "consumer.fulfillment.enqueue.error",
				"order_sn", event.OrderSN,
				"shop_id", event.ShopID,
				"error", err,
			)
		} else {
			logger.InfoContext(ctx, "fulfillment enqueued",
				"event", "consumer.fulfillment.enqueued",
				"order_sn", event.OrderSN,
				"shop_id", event.ShopID,
			)
		}
	}
}
