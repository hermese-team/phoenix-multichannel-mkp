package consumer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"

	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	"github.com/okdev/marketplace-sync/pkg/logger"
)

const (
	lifecycleConsumerGroup = "shopee-order-lifecycle"
	receivedInputTopic     = "order.received.v1.dev"
	lifecycleOutputTopic   = "order.lifecycle.v1.dev"
)

// LifecycleEvent is the append-only history record published to order.lifecycle.v1.dev.
//
// Every status transition produces one LifecycleEvent. Consumers can replay
// this topic to reconstruct the full order history or build audit trails.
//
// Key difference from ReceivedEvent (order.received.v1.dev):
//   - ReceivedEvent: signals "ingestion succeeded" to immediate downstream.
//   - LifecycleEvent: permanent history entry; never overwritten, only appended.
type LifecycleEvent struct {
	Channel    string `json:"channel"`
	EventType  string `json:"event_type"`  // canonical type (ORDER_CREATED, ORDER_SHIPPED, …)
	ShopID     int64  `json:"shop_id"`
	OrderSN    string `json:"order_sn"`
	OrderID    int64  `json:"order_id"`    // internal DB id assigned at first ingestion
	Status     string `json:"status"`
	IsNew      bool   `json:"is_new"`      // true = first time this order was ingested
	OccurredAt string `json:"occurred_at"` // RFC3339 — wall-clock time of ingestion
}

// LifecycleConsumer reads order.received.v1.dev and publishes an append-only
// LifecycleEvent to order.lifecycle.v1.dev for every successful ingestion.
//
// It is intentionally stateless: no DB writes, no dedup. The Kafka topic itself
// is the durable audit log; compact-and-retain policy preserves full history.
type LifecycleConsumer struct {
	cfg      kafkaAdapter.Config
	producer *kafkaAdapter.Producer
}

func NewLifecycleConsumer(cfg kafkaAdapter.Config, producer *kafkaAdapter.Producer) *LifecycleConsumer {
	return &LifecycleConsumer{cfg: cfg, producer: producer}
}

func (c *LifecycleConsumer) Start(ctx context.Context) error {
	kafkaClient, err := kgo.NewClient(
		kgo.SeedBrokers(c.cfg.Brokers...),
		kgo.ConsumerGroup(lifecycleConsumerGroup),
		kgo.ConsumeTopics(receivedInputTopic),
	)
	if err != nil {
		return err
	}
	defer kafkaClient.Close()

	logger.InfoContext(ctx, "lifecycle consumer started",
		"event", "lifecycle.started",
		"input_topic", receivedInputTopic,
		"output_topic", lifecycleOutputTopic,
		"group", lifecycleConsumerGroup,
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
				logger.ErrorContext(ctx, "lifecycle consumer fetch error",
					"event", "lifecycle.fetch.error",
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
			c.record(ctx, rec.Value)
		})
	}
}

func (c *LifecycleConsumer) record(ctx context.Context, data []byte) {
	var received ReceivedEvent
	if err := json.Unmarshal(data, &received); err != nil {
		logger.ErrorContext(ctx, "lifecycle unmarshal failed",
			"event", "lifecycle.unmarshal.error",
			"error", err,
		)
		return
	}

	evt := LifecycleEvent{
		Channel:    received.Channel,
		EventType:  received.EventType,
		ShopID:     received.ShopID,
		OrderSN:    received.OrderSN,
		OrderID:    received.OrderID,
		Status:     received.Status,
		IsNew:      received.IsNew,
		OccurredAt: time.Now().UTC().Format(time.RFC3339),
	}
	msgBytes, _ := json.Marshal(evt)

	if err := c.producer.Publish(ctx, lifecycleOutputTopic, []byte(received.OrderSN), msgBytes); err != nil {
		logger.ErrorContext(ctx, "lifecycle publish failed",
			"event", "lifecycle.publish_error",
			"order_sn", received.OrderSN,
			"error", err,
		)
		return
	}

	logger.InfoContext(ctx, "lifecycle event recorded",
		"event", "lifecycle.recorded",
		"order_sn", received.OrderSN,
		"shop_id", received.ShopID,
		"order_id", received.OrderID,
		"event_type", received.EventType,
		"status", received.Status,
		"is_new", received.IsNew,
	)
}
