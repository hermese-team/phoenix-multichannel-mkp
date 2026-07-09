package classifier

import (
	"context"
	"encoding/json"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"

	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	"github.com/okdev/marketplace-sync/pkg/logger"
)

const (
	classifierGroup = "shopee-classifier"
	inputTopic      = "order.raw.accepted.v1"
	OutputTopic     = "order.ingest.shopee.v1"
)

// rawEvent mirrors server.OrderRawEvent — separate to avoid circular imports.
type rawEvent struct {
	ShopID    int64  `json:"shop_id"`
	OrderSN   string `json:"order_sn"`
	Status    string `json:"status"`
	Code      int    `json:"code"`
	Timestamp int64  `json:"timestamp"`
}

// IngestEvent is published to order.ingest.shopee.v1.
// It carries a canonical EventType so downstream consumers are decoupled from Shopee push codes.
type IngestEvent struct {
	Channel   string `json:"channel"`
	EventType string `json:"event_type"`
	ShopID    int64  `json:"shop_id"`
	OrderSN   string `json:"order_sn"`
	Status    string `json:"status"`
	Code      int    `json:"code"`
	Timestamp int64  `json:"timestamp"`
}

// Consumer reads shopee.order.raw, classifies each event, and publishes
// a canonical IngestEvent to order.ingest.shopee.v1.
type Consumer struct {
	cfg      kafkaAdapter.Config
	producer *kafkaAdapter.Producer
}

func NewConsumer(cfg kafkaAdapter.Config, producer *kafkaAdapter.Producer) *Consumer {
	return &Consumer{cfg: cfg, producer: producer}
}

func (c *Consumer) Start(ctx context.Context) error {
	kafkaClient, err := kgo.NewClient(
		kgo.SeedBrokers(c.cfg.Brokers...),
		kgo.ConsumerGroup(classifierGroup),
		kgo.ConsumeTopics(inputTopic),
	)
	if err != nil {
		return err
	}
	defer kafkaClient.Close()

	logger.InfoContext(ctx, "classifier consumer started",
		"event", "classifier.started",
		"group", classifierGroup,
		"input_topic", inputTopic,
		"output_topic", OutputTopic,
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
				logger.ErrorContext(ctx, "classifier kafka fetch error",
					"event", "classifier.fetch.error",
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
			c.classify(ctx, rec.Value)
		})
	}
}

func (c *Consumer) classify(ctx context.Context, data []byte) {
	var raw rawEvent
	if err := json.Unmarshal(data, &raw); err != nil {
		logger.ErrorContext(ctx, "classifier unmarshal failed",
			"event", "classifier.unmarshal.error",
			"error", err,
		)
		return
	}

	eventType := Classify(raw.Code, raw.Status)

	out := IngestEvent{
		Channel:   "shopee",
		EventType: eventType,
		ShopID:    raw.ShopID,
		OrderSN:   raw.OrderSN,
		Status:    raw.Status,
		Code:      raw.Code,
		Timestamp: raw.Timestamp,
	}
	msgBytes, _ := json.Marshal(out)

	if err := c.producer.Publish(ctx, OutputTopic, []byte(raw.OrderSN), msgBytes); err != nil {
		logger.ErrorContext(ctx, "classifier publish failed",
			"event", "classifier.publish.error",
			"order_sn", raw.OrderSN,
			"shop_id", raw.ShopID,
			"error", err,
		)
		return
	}

	logger.InfoContext(ctx, "event classified",
		"event", "classifier.classified",
		"order_sn", raw.OrderSN,
		"shop_id", raw.ShopID,
		"code", raw.Code,
		"status", raw.Status,
		"event_type", eventType,
		"output_topic", OutputTopic,
	)
}
