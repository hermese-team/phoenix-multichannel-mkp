package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"git.amaze-x.com/phoenix/channel-adapter-shopee/config"
	"git.amaze-x.com/phoenix/channel-adapter-shopee/internal/infrastructure"
	"git.amaze-x.com/phoenix/channel-adapter-shopee/internal/shopee"
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

const (
	// Redis key for overlap-safe cursor: persists across restarts.
	redisCursorKey = "shopee:poller:cursor:%d" // :shop_id

	// Kafka record key: partition by shop_id + order_sn for ordering per order.
	kafkaKeyFmt = "%d:%s" // shop_id:order_sn
)

// Poller pulls orders from Shopee using cursor-based pagination and produces
// raw payloads to Kafka topic raw.order.accepted.v1.
//
// Architecture rules observed:
//   - Rate limited to 80 req/min via distributed Redis token bucket (20 reserved).
//   - Cursor is persisted in Redis so restarts don't re-fetch old pages.
//   - Overlap-safe: time window always includes a 5-minute back-overlap.
//   - Each order SN batch fetches detail (≤50 per call) then produces to Kafka.
//   - Never waits for PostgreSQL or downstream consumers — Kafka is the boundary.
type Poller struct {
	cfg      config.Config
	orderAPI *shopee.OrderAPI
	producer *infrastructure.KafkaProducer
	rdb      redis.UniversalClient
	limiter  *redis_rate.Limiter
	log      *zap.Logger
}

func NewPoller(
	cfg config.Config,
	orderAPI *shopee.OrderAPI,
	producer *infrastructure.KafkaProducer,
	rdb redis.UniversalClient,
	log *zap.Logger,
) *Poller {
	return &Poller{
		cfg:      cfg,
		orderAPI: orderAPI,
		producer: producer,
		rdb:      rdb,
		limiter:  redis_rate.NewLimiter(rdb),
		log:      log,
	}
}

// Run starts the polling loop. Blocks until ctx is cancelled.
func (p *Poller) Run(ctx context.Context) error {
	p.log.Info("poller starting",
		zap.Int64("shop_id", p.cfg.Shopee.ShopID),
		zap.Int("interval_sec", p.cfg.Poller.IntervalSec),
		zap.Int("rate_limit_per_min", p.cfg.Poller.RateLimitPerMin),
	)

	ticker := time.NewTicker(time.Duration(p.cfg.Poller.IntervalSec) * time.Second)
	defer ticker.Stop()

	// Run immediately on startup, then on every tick.
	if err := p.poll(ctx); err != nil {
		p.log.Error("poll cycle failed", zap.Error(err))
	}

	for {
		select {
		case <-ctx.Done():
			p.log.Info("poller stopped")
			return ctx.Err()
		case <-ticker.C:
			if err := p.poll(ctx); err != nil {
				p.log.Error("poll cycle failed", zap.Error(err))
				// Non-fatal: log and continue. Next tick will retry.
			}
		}
	}
}

// poll fetches all pages of updated orders since the last cursor/timestamp.
func (p *Poller) poll(ctx context.Context) error {
	ctx, span := otel.Tracer("shopee.poller").Start(ctx, "poller.cycle")
	defer span.End()

	now := time.Now().Unix()
	timeTo := now
	timeFrom := now - p.cfg.Poller.LookbackSec

	// Load persisted cursor — overlap-safe, no missed orders across restarts.
	cursor, err := p.loadCursor(ctx)
	if err != nil {
		p.log.Warn("failed to load cursor, starting from scratch", zap.Error(err))
		cursor = ""
	}

	var totalFetched int
	for {
		// Rate limit: 80 req/min operating budget, Shopee quota is 100/min.
		if err := p.rateWait(ctx); err != nil {
			return fmt.Errorf("rate limiter: %w", err)
		}

		briefs, nextCursor, more, err := p.orderAPI.GetOrderList(ctx, shopee.GetOrderListRequest{
			TimeRangeField: "update_time",
			TimeFrom:       timeFrom,
			TimeTo:         timeTo,
			PageSize:       p.cfg.Poller.PageSize,
			Cursor:         cursor,
		})
		if err != nil {
			span.RecordError(err)
			return fmt.Errorf("get_order_list: %w", err)
		}

		if len(briefs) > 0 {
			sns := make([]string, len(briefs))
			for i, b := range briefs {
				sns[i] = b.OrderSN
			}

			// Fetch full detail — 50 SNs per call, auto-chunked.
			if err := p.rateWait(ctx); err != nil {
				return err
			}
			orders, err := p.orderAPI.GetOrderDetailAll(ctx, sns,
				"buyer_username", "recipient_address", "item_list",
			)
			if err != nil {
				span.RecordError(err)
				return fmt.Errorf("get_order_detail: %w", err)
			}

			for _, order := range orders {
				if err := p.produceOrder(ctx, order); err != nil {
					return fmt.Errorf("producing order %s: %w", order.OrderSN, err)
				}
			}
			totalFetched += len(orders)
		}

		// Persist cursor so next poll (or restart) continues from here.
		if nextCursor != "" {
			if err := p.saveCursor(ctx, nextCursor); err != nil {
				p.log.Warn("failed to save cursor", zap.Error(err))
			}
		}

		if !more {
			break
		}
		cursor = nextCursor
	}

	span.SetAttributes(attribute.Int("poller.orders_fetched", totalFetched))
	p.log.Info("poll cycle complete", zap.Int("fetched", totalFetched))
	return nil
}

// produceOrder serialises a full order and produces it to Kafka.
func (p *Poller) produceOrder(ctx context.Context, order shopee.Order) error {
	payload, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("marshalling order %s: %w", order.OrderSN, err)
	}

	key := fmt.Sprintf(kafkaKeyFmt, p.cfg.Shopee.ShopID, order.OrderSN)
	topic := p.cfg.Kafka.TopicRawOrderAccepted

	return p.producer.Produce(ctx, topic, key, payload)
}

// rateWait blocks until a token is available in the distributed rate bucket.
// Uses go-redis/redis_rate GCRA limiter keyed per shop_id.
func (p *Poller) rateWait(ctx context.Context) error {
	key := fmt.Sprintf("shopee:ratelimit:%d", p.cfg.Shopee.ShopID)
	limit := redis_rate.PerMinute(p.cfg.Poller.RateLimitPerMin)

	for {
		res, err := p.limiter.Allow(ctx, key, limit)
		if err != nil {
			return fmt.Errorf("rate limiter allow: %w", err)
		}
		if res.Allowed > 0 {
			return nil
		}
		// Wait for the retry-after duration suggested by the limiter.
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(res.RetryAfter):
		}
	}
}

func (p *Poller) loadCursor(ctx context.Context) (string, error) {
	key := fmt.Sprintf(redisCursorKey, p.cfg.Shopee.ShopID)
	val, err := p.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func (p *Poller) saveCursor(ctx context.Context, cursor string) error {
	key := fmt.Sprintf(redisCursorKey, p.cfg.Shopee.ShopID)
	return p.rdb.Set(ctx, key, cursor, 24*time.Hour).Err()
}
