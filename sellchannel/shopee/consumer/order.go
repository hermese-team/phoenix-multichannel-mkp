package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"

	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	pgAdapter     "github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/client"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/tokenstore"
)

const (
	orderConsumerGroup = "shopee-order-enricher"
	orderIngestTopic   = "order.ingest.shopee.v1"
	orderEnrichedTopic = "order.enriched.v1"
	orderDLQTopic      = "order.enriched.shopee.dlq"

	maxBatchSize   = 50              // Shopee API limit per GetOrderDetail call
	coalesceWindow = 2 * time.Second // buffer window: merge duplicate order_sn pushes
	maxRetries     = 3
	apiRatePerSec  = 5.0 // conservative Shopee API rate limit (calls/sec)
)

// orderIngestEvent mirrors classifier.IngestEvent — defined here to avoid circular imports.
type orderIngestEvent struct {
	Channel   string `json:"channel"`
	EventType string `json:"event_type"`
	ShopID    int64  `json:"shop_id"`
	OrderSN   string `json:"order_sn"`
	Status    string `json:"status"`
	Code      int    `json:"code"`
	Timestamp int64  `json:"timestamp"`
}

// EnrichedEvent is published to order.enriched.v1.
type EnrichedEvent struct {
	Channel    string             `json:"channel"`
	EventType  string             `json:"event_type"`
	ShopID     int64              `json:"shop_id"`
	OrderSN    string             `json:"order_sn"`
	Order      client.OrderDetail `json:"order"`
	EnrichedAt string             `json:"enriched_at"`
}

// DLQEvent is published to order.enriched.shopee.dlq after maxRetries failures.
type DLQEvent struct {
	Channel   string `json:"channel"`
	EventType string `json:"event_type"`
	ShopID    int64  `json:"shop_id"`
	OrderSN   string `json:"order_sn"`
	Error     string `json:"error"`
	Attempts  int    `json:"attempts"`
	FailedAt  string `json:"failed_at"`
}

// OrderConsumer reads order.ingest.shopee.v1, coalesces notifications within a
// time window, bulk-fetches order detail from Shopee API, and publishes to
// order.enriched.v1. Permanent failures go to the DLQ topic.
type OrderConsumer struct {
	cfg             kafkaAdapter.Config
	shopeeClient    *client.Client
	tokens          *tokenstore.Store
	producer        *kafkaAdapter.Producer
	fulfillmentRepo *pgAdapter.ShopeeFulfillmentRepository
	rateTicker      *time.Ticker // rate limiter: 1 tick per API slot
}

func NewOrderConsumer(cfg kafkaAdapter.Config, shopeeClient *client.Client, tokens *tokenstore.Store, producer *kafkaAdapter.Producer, fulfillmentRepo *pgAdapter.ShopeeFulfillmentRepository) *OrderConsumer {
	interval := time.Duration(float64(time.Second) / apiRatePerSec)
	return &OrderConsumer{
		cfg:             cfg,
		shopeeClient:    shopeeClient,
		tokens:          tokens,
		producer:        producer,
		fulfillmentRepo: fulfillmentRepo,
		rateTicker:      time.NewTicker(interval),
	}
}

func (c *OrderConsumer) Start(ctx context.Context) error {
	kafkaClient, err := kgo.NewClient(
		kgo.SeedBrokers(c.cfg.Brokers...),
		kgo.ConsumerGroup(orderConsumerGroup),
		kgo.ConsumeTopics(orderIngestTopic),
	)
	if err != nil {
		return err
	}
	defer kafkaClient.Close()

	logger.InfoContext(ctx, "order consumer started",
		"event", "consumer.started",
		"topic", orderIngestTopic,
		"group", orderConsumerGroup,
		"coalesce_window_ms", coalesceWindow.Milliseconds(),
		"max_batch_size", maxBatchSize,
	)

	// buffer: order_sn → latest event (coalesce duplicates)
	buffer := make(map[string]orderIngestEvent)
	recordsCh := make(chan orderIngestEvent, 1000)

	coalesceTicker := time.NewTicker(coalesceWindow)
	defer coalesceTicker.Stop()

	// Kafka reader goroutine — feeds recordsCh
	go func() {
		for {
			if ctx.Err() != nil {
				return
			}
			fetches := kafkaClient.PollFetches(ctx)
			if errs := fetches.Errors(); len(errs) > 0 {
				if ctx.Err() != nil {
					return
				}
				for _, e := range errs {
					logger.ErrorContext(ctx, "kafka fetch error",
						"event", "consumer.fetch.error",
						"error", e.Err,
					)
				}
				select {
				case <-ctx.Done():
					return
				case <-time.After(time.Second):
				}
				continue
			}
			fetches.EachRecord(func(rec *kgo.Record) {
				var event orderIngestEvent
				if err := json.Unmarshal(rec.Value, &event); err != nil {
					logger.ErrorContext(ctx, "unmarshal ingest event failed",
						"event", "consumer.unmarshal.error",
						"error", err,
					)
					return
				}
				select {
				case recordsCh <- event:
				case <-ctx.Done():
				}
			})
		}
	}()

	flush := func() {
		if len(buffer) == 0 {
			return
		}
		batch := buffer
		buffer = make(map[string]orderIngestEvent)
		logger.InfoContext(ctx, "flushing enrichment batch",
			"event", "consumer.batch.flush",
			"size", len(batch),
		)
		c.processBatch(ctx, batch)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case event := <-recordsCh:
			prev, exists := buffer[event.OrderSN]
			buffer[event.OrderSN] = event // last write wins
			if exists {
				logger.DebugContext(ctx, "coalesced duplicate notification",
					"event", "consumer.coalesce",
					"order_sn", event.OrderSN,
					"prev_status", prev.Status,
					"new_status", event.Status,
				)
			}
			if len(buffer) >= maxBatchSize {
				flush()
				coalesceTicker.Reset(coalesceWindow)
			}
		case <-coalesceTicker.C:
			flush()
		}
	}
}

// processBatch groups events by shop, rate-limits API calls, and publishes results.
func (c *OrderConsumer) processBatch(ctx context.Context, batch map[string]orderIngestEvent) {
	// Group by shop_id — each shop has its own token.
	byShop := make(map[int64][]orderIngestEvent)
	for _, event := range batch {
		byShop[event.ShopID] = append(byShop[event.ShopID], event)
	}

	for shopID, events := range byShop {
		accessToken, _, err := c.tokens.Get(ctx, shopID)
		if err != nil || accessToken == "" {
			logger.ErrorContext(ctx, "get token failed for batch",
				"event", "consumer.token.error",
				"shop_id", shopID,
				"error", err,
			)
			for _, e := range events {
				c.publishDLQ(ctx, e, fmt.Errorf("no token: %w", err), 0)
			}
			continue
		}

		// Chunk into groups of maxBatchSize.
		for i := 0; i < len(events); i += maxBatchSize {
			end := i + maxBatchSize
			if end > len(events) {
				end = len(events)
			}
			chunk := events[i:end]
			c.fetchAndPublish(ctx, shopID, accessToken, chunk)
		}
	}
}

func (c *OrderConsumer) fetchAndPublish(ctx context.Context, shopID int64, accessToken string, chunk []orderIngestEvent) {
	sns := make([]string, len(chunk))
	for i, e := range chunk {
		sns[i] = e.OrderSN
	}

	// Rate-limit: wait for the next API slot.
	select {
	case <-ctx.Done():
		return
	case <-c.rateTicker.C:
	}

	var orders []client.OrderDetail
	var lastErr error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		orders, lastErr = c.shopeeClient.GetOrderDetail(ctx, shopID, accessToken, sns)
		if lastErr == nil {
			break
		}
		logger.WarnContext(ctx, "get order detail failed, retrying",
			"event", "consumer.get_order_detail.retry",
			"shop_id", shopID,
			"attempt", attempt,
			"max_retries", maxRetries,
			"error", lastErr,
		)
		if attempt < maxRetries {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Duration(attempt) * time.Second):
			}
		}
	}

	if lastErr != nil {
		logger.ErrorContext(ctx, "get order detail failed after retries, sending to DLQ",
			"event", "consumer.get_order_detail.dlq",
			"shop_id", shopID,
			"order_sns", sns,
			"error", lastErr,
		)
		for _, e := range chunk {
			c.publishDLQ(ctx, e, lastErr, maxRetries)
		}
		return
	}

	// Index results by order_sn.
	orderMap := make(map[string]client.OrderDetail, len(orders))
	for _, o := range orders {
		o.ShopID = shopID // Shopee API omits shop_id in response
		orderMap[o.OrderSN] = o
	}

	for _, e := range chunk {
		order, ok := orderMap[e.OrderSN]
		if !ok {
			c.publishDLQ(ctx, e, fmt.Errorf("order not found in GetOrderDetail response"), maxRetries)
			continue
		}
		c.publishEnriched(ctx, e, order)
	}
}

func (c *OrderConsumer) publishEnriched(ctx context.Context, e orderIngestEvent, order client.OrderDetail) {
	out := EnrichedEvent{
		Channel:    "shopee",
		EventType:  e.EventType,
		ShopID:     e.ShopID,
		OrderSN:    e.OrderSN,
		Order:      order,
		EnrichedAt: time.Now().UTC().Format(time.RFC3339),
	}
	msgBytes, _ := json.Marshal(out)
	if err := c.producer.Publish(ctx, orderEnrichedTopic, []byte(e.OrderSN), msgBytes); err != nil {
		logger.ErrorContext(ctx, "publish enriched event failed",
			"event", "consumer.enriched.publish_error",
			"order_sn", e.OrderSN,
			"error", err,
		)
		return
	}

	logger.InfoContext(ctx, "order enriched",
		"event", "consumer.order.enriched",
		"order_sn", e.OrderSN,
		"shop_id", e.ShopID,
		"event_type", e.EventType,
		"order_status", order.OrderStatus,
		"topic", orderEnrichedTopic,
	)

	// Enqueue fulfillment when order is ready to ship.
	if order.OrderStatus == "READY_TO_SHIP" && c.fulfillmentRepo != nil {
		if err := c.fulfillmentRepo.Enqueue(ctx, e.ShopID, e.OrderSN,
			pgAdapter.ShopeeDeliveryShopeeLogistics, ""); err != nil {
			logger.ErrorContext(ctx, "enqueue fulfillment failed",
				"event", "consumer.fulfillment.enqueue.error",
				"order_sn", e.OrderSN,
				"shop_id", e.ShopID,
				"error", err,
			)
		} else {
			logger.InfoContext(ctx, "fulfillment enqueued",
				"event", "consumer.fulfillment.enqueued",
				"order_sn", e.OrderSN,
				"shop_id", e.ShopID,
			)
		}
	}
}

func (c *OrderConsumer) publishDLQ(ctx context.Context, e orderIngestEvent, cause error, attempts int) {
	dlq := DLQEvent{
		Channel:   "shopee",
		EventType: e.EventType,
		ShopID:    e.ShopID,
		OrderSN:   e.OrderSN,
		Error:     cause.Error(),
		Attempts:  attempts,
		FailedAt:  time.Now().UTC().Format(time.RFC3339),
	}
	msgBytes, _ := json.Marshal(dlq)
	if err := c.producer.Publish(ctx, orderDLQTopic, []byte(e.OrderSN), msgBytes); err != nil {
		logger.ErrorContext(ctx, "publish DLQ failed",
			"event", "consumer.dlq.publish_error",
			"order_sn", e.OrderSN,
			"error", err,
		)
		return
	}
	logger.WarnContext(ctx, "order sent to DLQ",
		"event", "consumer.order.dlq",
		"order_sn", e.OrderSN,
		"shop_id", e.ShopID,
		"cause", cause.Error(),
		"attempts", attempts,
		"topic", orderDLQTopic,
	)
}
