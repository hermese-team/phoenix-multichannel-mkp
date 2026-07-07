package consumer

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"

	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	pgAdapter     "github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
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

	log.Println("shopee order consumer started")
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
				log.Printf("[order-consumer] fetch error: %v", e.Err)
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
				log.Printf("[order-consumer] unmarshal: %v", err)
				return
			}
			c.process(ctx, event)
		})
	}
}

func (c *OrderConsumer) process(ctx context.Context, event orderRawEvent) {
	accessToken, _, err := c.tokens.Get(ctx, event.ShopID)
	if err != nil {
		log.Printf("[order-consumer] get token shop %d: %v", event.ShopID, err)
		return
	}
	if accessToken == "" {
		log.Printf("[order-consumer] shop %d not authorized", event.ShopID)
		return
	}

	orders, err := c.shopeeClient.GetOrderDetail(ctx, event.ShopID, accessToken, []string{event.OrderSN})
	if err != nil {
		log.Printf("[order-consumer] get order detail %s: %v", event.OrderSN, err)
		return
	}
	if len(orders) == 0 {
		log.Printf("[order-consumer] order %s not found", event.OrderSN)
		return
	}

	order := orders[0]
	order.ShopID = event.ShopID // Shopee API omits shop_id in response

	msgBytes, _ := json.Marshal(order)
	if err := c.producer.Publish(ctx, orderDetailTopic, []byte(event.OrderSN), msgBytes); err != nil {
		log.Printf("[order-consumer] publish %s: %v", event.OrderSN, err)
		return
	}

	log.Printf("[order-consumer] published order %s (status=%s)", event.OrderSN, event.Status)

	// Enqueue fulfillment when order is ready to ship.
	if order.OrderStatus == "READY_TO_SHIP" && c.fulfillmentRepo != nil {
		if err := c.fulfillmentRepo.Enqueue(ctx, event.ShopID, event.OrderSN,
			pgAdapter.ShopeeDeliveryShopeeLogistics, ""); err != nil {
			log.Printf("[order-consumer] enqueue fulfillment %s: %v", event.OrderSN, err)
		} else {
			log.Printf("[order-consumer] enqueued fulfillment %s", event.OrderSN)
		}
	}
}
