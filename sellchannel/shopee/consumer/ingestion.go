package consumer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"

	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	pgAdapter     "github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	orderDomain   "github.com/okdev/marketplace-sync/internal/domain/order"
	"github.com/okdev/marketplace-sync/pkg/logger"
)

const (
	ingestionConsumerGroup = "shopee-order-ingestion"
	enrichedInputTopic     = "order.enriched.v1"
	receivedTopic          = "order.received.v1"

	ingestionBatchSize = 100
	ingestionWindow    = 500 * time.Millisecond
)

// ReceivedEvent is published to order.received.v1 after successful ingestion.
type ReceivedEvent struct {
	Channel    string `json:"channel"`
	EventType  string `json:"event_type"`
	ShopID     int64  `json:"shop_id"`
	OrderSN    string `json:"order_sn"`
	OrderID    int64  `json:"order_id"`    // internal DB id
	Status     string `json:"status"`
	IsNew      bool   `json:"is_new"`      // true = first time we saw this order
	ReceivedAt string `json:"received_at"`
}

// IngestionConsumer reads order.enriched.v1, micro-batches the events,
// upserts them into the orders table (idempotent), and publishes
// order.received.v1 for each successful write.
type IngestionConsumer struct {
	cfg       kafkaAdapter.Config
	orderRepo *pgAdapter.OrderRepository
	producer  *kafkaAdapter.Producer
}

func NewIngestionConsumer(cfg kafkaAdapter.Config, orderRepo *pgAdapter.OrderRepository, producer *kafkaAdapter.Producer) *IngestionConsumer {
	return &IngestionConsumer{cfg: cfg, orderRepo: orderRepo, producer: producer}
}

func (c *IngestionConsumer) Start(ctx context.Context) error {
	kafkaClient, err := kgo.NewClient(
		kgo.SeedBrokers(c.cfg.Brokers...),
		kgo.ConsumerGroup(ingestionConsumerGroup),
		kgo.ConsumeTopics(enrichedInputTopic),
	)
	if err != nil {
		return err
	}
	defer kafkaClient.Close()

	logger.InfoContext(ctx, "ingestion consumer started",
		"event", "ingestion.started",
		"topic", enrichedInputTopic,
		"group", ingestionConsumerGroup,
		"batch_size", ingestionBatchSize,
		"window_ms", ingestionWindow.Milliseconds(),
	)

	recordsCh := make(chan EnrichedEvent, 1000)
	buffer := make([]EnrichedEvent, 0, ingestionBatchSize)
	windowTicker := time.NewTicker(ingestionWindow)
	defer windowTicker.Stop()

	// Kafka reader goroutine
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
					logger.ErrorContext(ctx, "ingestion kafka fetch error",
						"event", "ingestion.fetch.error",
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
				var event EnrichedEvent
				if err := json.Unmarshal(rec.Value, &event); err != nil {
					logger.ErrorContext(ctx, "ingestion unmarshal failed",
						"event", "ingestion.unmarshal.error",
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
		buffer = make([]EnrichedEvent, 0, ingestionBatchSize)
		c.ingestBatch(ctx, batch)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case event := <-recordsCh:
			buffer = append(buffer, event)
			if len(buffer) >= ingestionBatchSize {
				flush()
				windowTicker.Reset(ingestionWindow)
			}
		case <-windowTicker.C:
			flush()
		}
	}
}

func (c *IngestionConsumer) ingestBatch(ctx context.Context, batch []EnrichedEvent) {
	logger.InfoContext(ctx, "ingesting order batch",
		"event", "ingestion.batch",
		"size", len(batch),
	)

	for _, event := range batch {
		o := &orderDomain.Order{
			ShopID:          event.ShopID,
			SellChannelID:   event.OrderSN,
			SellChannelType: "shopee",
			Status:          orderDomain.Status(event.Order.OrderStatus),
			TotalAmount:     event.Order.TotalAmount,
		}

		isNew, err := c.orderRepo.Upsert(ctx, o)
		if err != nil {
			logger.ErrorContext(ctx, "order upsert failed",
				"event", "ingestion.upsert.error",
				"order_sn", event.OrderSN,
				"shop_id", event.ShopID,
				"error", err,
			)
			continue
		}

		logger.InfoContext(ctx, "order ingested",
			"event", "ingestion.order.ingested",
			"order_sn", event.OrderSN,
			"shop_id", event.ShopID,
			"order_id", o.ID,
			"status", o.Status,
			"is_new", isNew,
		)

		received := ReceivedEvent{
			Channel:    "shopee",
			EventType:  event.EventType,
			ShopID:     event.ShopID,
			OrderSN:    event.OrderSN,
			OrderID:    o.ID,
			Status:     string(o.Status),
			IsNew:      isNew,
			ReceivedAt: time.Now().UTC().Format(time.RFC3339),
		}
		msgBytes, _ := json.Marshal(received)
		if err := c.producer.Publish(ctx, receivedTopic, []byte(event.OrderSN), msgBytes); err != nil {
			logger.ErrorContext(ctx, "publish order.received.v1 failed",
				"event", "ingestion.received.publish_error",
				"order_sn", event.OrderSN,
				"error", err,
			)
			continue
		}

		logger.InfoContext(ctx, "order.received.v1 published",
			"event", "ingestion.received.published",
			"order_sn", event.OrderSN,
			"shop_id", event.ShopID,
			"order_id", o.ID,
			"is_new", isNew,
			"topic", receivedTopic,
		)
	}
}
