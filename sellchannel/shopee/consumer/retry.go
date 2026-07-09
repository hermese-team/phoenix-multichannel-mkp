package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"

	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/client"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/tokenstore"
)

const (
	retryConsumerGroup = "shopee-order-retry"
	orderRetryTopic    = "order.retry.v1.dev"
	maxRetryAttempts   = 5
)

// retryBackoff[i] is the wait duration before attempt i+1 (0-indexed).
// attempt 1 → 30s, 2 → 2m, 3 → 10m, 4 → 30m, 5 → DLQ
var retryBackoff = []time.Duration{
	30 * time.Second,
	2 * time.Minute,
	10 * time.Minute,
	30 * time.Minute,
}

func backoffFor(attempt int) time.Duration {
	idx := attempt - 1
	if idx < 0 {
		idx = 0
	}
	if idx >= len(retryBackoff) {
		idx = len(retryBackoff) - 1
	}
	return retryBackoff[idx]
}

// RetryEvent is published to order.retry.v1.dev when GetOrderDetail fails.
// The RetryConsumer re-attempts enrichment after NextAttemptAt.
type RetryEvent struct {
	Channel       string           `json:"channel"`
	EventType     string           `json:"event_type"`
	ShopID        int64            `json:"shop_id"`
	OrderSN       string           `json:"order_sn"`
	Attempt       int              `json:"attempt"`         // current attempt number (1-based)
	MaxAttempts   int              `json:"max_attempts"`
	NextAttemptAt time.Time        `json:"next_attempt_at"` // wall-clock time, not a delay
	LastError     string           `json:"last_error"`
	Original      orderIngestEvent `json:"original_event"`  // preserved for re-processing
}

// RetryConsumer reads order.retry.v1.dev, waits for the scheduled backoff,
// re-attempts GetOrderDetail, and routes to order.enriched.v1.dev on success
// or order.dlq.v1.dev when MaxAttempts is exhausted.
//
// Backoff schedule (attempt → delay before next attempt):
//
//	1 → 30s  |  2 → 2m  |  3 → 10m  |  4 → 30m  |  5 → DLQ
type RetryConsumer struct {
	cfg          kafkaAdapter.Config
	shopeeClient *client.Client
	tokens       *tokenstore.Store
	producer     *kafkaAdapter.Producer
	rateTicker   *time.Ticker
}

func NewRetryConsumer(
	cfg kafkaAdapter.Config,
	shopeeClient *client.Client,
	tokens *tokenstore.Store,
	producer *kafkaAdapter.Producer,
) *RetryConsumer {
	interval := time.Duration(float64(time.Second) / apiRatePerSec)
	return &RetryConsumer{
		cfg:          cfg,
		shopeeClient: shopeeClient,
		tokens:       tokens,
		producer:     producer,
		rateTicker:   time.NewTicker(interval),
	}
}

func (c *RetryConsumer) Start(ctx context.Context) error {
	kafkaClient, err := kgo.NewClient(
		kgo.SeedBrokers(c.cfg.Brokers...),
		kgo.ConsumerGroup(retryConsumerGroup),
		kgo.ConsumeTopics(orderRetryTopic),
	)
	if err != nil {
		return err
	}
	defer kafkaClient.Close()

	logger.InfoContext(ctx, "retry consumer started",
		"event", "retry.started",
		"topic", orderRetryTopic,
		"group", retryConsumerGroup,
		"max_attempts", maxRetryAttempts,
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
				logger.ErrorContext(ctx, "retry consumer fetch error",
					"event", "retry.fetch.error",
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
			var evt RetryEvent
			if err := json.Unmarshal(rec.Value, &evt); err != nil {
				logger.ErrorContext(ctx, "retry unmarshal failed",
					"event", "retry.unmarshal.error",
					"error", err,
				)
				return
			}
			c.process(ctx, evt)
		})
	}
}

func (c *RetryConsumer) process(ctx context.Context, evt RetryEvent) {
	// Wait until NextAttemptAt before calling the API.
	if wait := time.Until(evt.NextAttemptAt); wait > 0 {
		logger.InfoContext(ctx, "retry waiting for backoff",
			"event", "retry.waiting",
			"order_sn", evt.OrderSN,
			"attempt", evt.Attempt,
			"wait_s", int(wait.Seconds()),
		)
		select {
		case <-ctx.Done():
			return
		case <-time.After(wait):
		}
	}

	accessToken, _, err := c.tokens.Get(ctx, evt.ShopID)
	if err != nil || accessToken == "" {
		c.handleFailure(ctx, evt, fmt.Errorf("no token: %w", err))
		return
	}

	// Rate-limit against Shopee API.
	select {
	case <-ctx.Done():
		return
	case <-c.rateTicker.C:
	}

	orders, err := c.shopeeClient.GetOrderDetail(ctx, evt.ShopID, accessToken, []string{evt.OrderSN})
	if err != nil {
		c.handleFailure(ctx, evt, err)
		return
	}
	if len(orders) == 0 {
		c.handleFailure(ctx, evt, fmt.Errorf("order not found in GetOrderDetail response"))
		return
	}

	order := orders[0]
	order.ShopID = evt.ShopID

	out := EnrichedEvent{
		Channel:    "shopee",
		EventType:  evt.Original.EventType,
		ShopID:     evt.ShopID,
		OrderSN:    evt.OrderSN,
		Order:      order,
		EnrichedAt: time.Now().UTC().Format(time.RFC3339),
	}
	msgBytes, _ := json.Marshal(out)
	if err := c.producer.Publish(ctx, orderEnrichedTopic, []byte(evt.OrderSN), msgBytes); err != nil {
		logger.ErrorContext(ctx, "retry publish enriched failed",
			"event", "retry.enriched.publish_error",
			"order_sn", evt.OrderSN,
			"error", err,
		)
		return
	}

	logger.InfoContext(ctx, "retry succeeded",
		"event", "retry.succeeded",
		"order_sn", evt.OrderSN,
		"shop_id", evt.ShopID,
		"attempt", evt.Attempt,
		"topic", orderEnrichedTopic,
	)
}

func (c *RetryConsumer) handleFailure(ctx context.Context, evt RetryEvent, cause error) {
	if evt.Attempt >= evt.MaxAttempts {
		// All attempts exhausted → DLQ.
		dlq := DLQEvent{
			Channel:   "shopee",
			EventType: evt.Original.EventType,
			ShopID:    evt.ShopID,
			OrderSN:   evt.OrderSN,
			Error:     cause.Error(),
			Attempts:  evt.Attempt,
			FailedAt:  time.Now().UTC().Format(time.RFC3339),
		}
		msgBytes, _ := json.Marshal(dlq)
		if err := c.producer.Publish(ctx, orderDLQTopic, []byte(evt.OrderSN), msgBytes); err != nil {
			logger.ErrorContext(ctx, "retry publish DLQ failed",
				"event", "retry.dlq.publish_error",
				"order_sn", evt.OrderSN,
				"error", err,
			)
			return
		}
		logger.WarnContext(ctx, "retry exhausted, sent to DLQ",
			"event", "retry.exhausted",
			"order_sn", evt.OrderSN,
			"shop_id", evt.ShopID,
			"attempts", evt.Attempt,
			"cause", cause.Error(),
		)
		return
	}

	// Schedule next attempt with exponential backoff.
	nextAttempt := evt.Attempt + 1
	delay := backoffFor(nextAttempt)
	next := RetryEvent{
		Channel:       evt.Channel,
		EventType:     evt.EventType,
		ShopID:        evt.ShopID,
		OrderSN:       evt.OrderSN,
		Attempt:       nextAttempt,
		MaxAttempts:   evt.MaxAttempts,
		NextAttemptAt: time.Now().Add(delay),
		LastError:     cause.Error(),
		Original:      evt.Original,
	}
	msgBytes, _ := json.Marshal(next)
	if err := c.producer.Publish(ctx, orderRetryTopic, []byte(evt.OrderSN), msgBytes); err != nil {
		logger.ErrorContext(ctx, "retry re-publish failed",
			"event", "retry.republish_error",
			"order_sn", evt.OrderSN,
			"error", err,
		)
		return
	}

	logger.WarnContext(ctx, "retry scheduled",
		"event", "retry.scheduled",
		"order_sn", evt.OrderSN,
		"shop_id", evt.ShopID,
		"attempt", nextAttempt,
		"max_attempts", evt.MaxAttempts,
		"next_attempt_at", next.NextAttemptAt.Format(time.RFC3339),
		"last_error", cause.Error(),
		"delay_s", int(delay.Seconds()),
	)
}
