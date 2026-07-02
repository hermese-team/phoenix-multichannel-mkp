package consumer

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"

	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/ordersync"
)

// Consumer reads Lazada webhook messages from Kafka and syncs each order's
// detail + items into Postgres via the shared syncer.
type Consumer struct {
	brokers []string
	group   string
	topic   string
	syncer  *ordersync.Syncer
}

func New(kafkaCfg kafkaAdapter.Config, syncer *ordersync.Syncer) *Consumer {
	return &Consumer{
		brokers: kafkaCfg.Brokers,
		group:   getEnv("LAZADA_WEBHOOK_GROUP", "lazada-webhook-consumer"),
		topic:   getEnv("LAZADA_WEBHOOK_TOPIC", "lazada.order.webhook"),
		syncer:  syncer,
	}
}

// message is the subset of the Lazada webhook payload the consumer needs.
type message struct {
	Data struct {
		TradeOrderID string `json:"trade_order_id"`
	} `json:"data"`
}

func (c *Consumer) Start(ctx context.Context) error {
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(c.brokers...),
		kgo.ConsumerGroup(c.group),
		kgo.ConsumeTopics(c.topic),
		kgo.DisableAutoCommit(), // commit only after a record is synced (at-least-once)
	)
	if err != nil {
		return err
	}
	defer cl.Close()

	logger.Info("lazada webhook consumer started", "topic", c.topic, "group", c.group)
	for {
		if ctx.Err() != nil {
			return nil
		}
		fetches := cl.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			if ctx.Err() != nil {
				return nil
			}
			for _, e := range errs {
				logger.Error("kafka fetch", "error", e.Err)
			}
			// Back off so a persistent fetch error (broker down, rebalance
			// loop, …) can't spin this loop at full CPU and flood the broker.
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(time.Second):
			}
			continue
		}
		fetches.EachRecord(func(rec *kgo.Record) {
			var m message
			if err := json.Unmarshal(rec.Value, &m); err != nil {
				logger.Error("unmarshal webhook message", "error", err)
				return
			}
			if m.Data.TradeOrderID == "" {
				logger.Warn("webhook message has no trade_order_id", "value", string(rec.Value))
				return
			}
			if err := c.syncer.SyncOrder(ctx, m.Data.TradeOrderID); err != nil {
				// Not committed → the record is re-read on restart/rebalance.
				logger.Error("sync order from webhook", "order", m.Data.TradeOrderID, "error", err)
				return
			}
			if err := cl.CommitRecords(ctx, rec); err != nil {
				logger.Error("commit offset", "order", m.Data.TradeOrderID, "error", err)
			}
		})
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
