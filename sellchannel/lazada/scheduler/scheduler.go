package scheduler

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/ordersync"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/tokenstore"
)

// Scheduler periodically pulls the Lazada order list and syncs each order via
// the shared ordersync.Syncer.
type Scheduler struct {
	cron   *cron.Cron
	syncer *ordersync.Syncer
	tokens *tokenstore.Store

	spec   string
	window time.Duration
	limit  int
}

func New(syncer *ordersync.Syncer, tokens *tokenstore.Store) *Scheduler {
	return &Scheduler{
		cron:   cron.New(),
		syncer: syncer,
		tokens: tokens,
		spec:   getEnv("LAZADA_SYNC_SPEC", "@every 5m"),
		window: getEnvDuration("LAZADA_SYNC_WINDOW", 24*time.Hour),
		limit:  getEnvInt("LAZADA_SYNC_LIMIT", 50),
	}
}

func (s *Scheduler) Start(ctx context.Context) error {
	// Tokens live in Redis. Require at least a refresh token to bootstrap;
	// seed it once with the lazada-token command.
	if _, refresh, err := s.tokens.Get(ctx); err != nil {
		return err
	} else if refresh == "" {
		return fmt.Errorf("no lazada refresh token in redis — seed it first: `make seed-lazada-token REFRESH=<token>`")
	}

	// Run once immediately, then on the cron schedule.
	s.runSync(ctx)

	if _, err := s.cron.AddFunc(s.spec, func() { s.runSync(context.Background()) }); err != nil {
		return err
	}
	s.cron.Start()
	logger.Info("lazada scheduler started", "spec", s.spec, "window", s.window.String(), "limit", s.limit)

	<-ctx.Done()
	s.cron.Stop()
	logger.Info("lazada scheduler stopped")
	return nil
}

func (s *Scheduler) runSync(ctx context.Context) {
	if err := s.syncOrders(ctx); err != nil {
		logger.Error("lazada order sync failed", "error", err)
	}
}

func (s *Scheduler) syncOrders(ctx context.Context) error {
	before := time.Now()
	after := before.Add(-s.window)

	processed, saved := 0, 0
	for offset := 0; ; {
		total, ids, err := s.syncer.ListOrders(ctx, after, before, offset)
		if err != nil {
			return err
		}
		if offset == 0 {
			logger.Info("lazada order list fetched", "total", total)
		}
		if len(ids) == 0 {
			break
		}
		for _, id := range ids {
			if s.limit > 0 && processed >= s.limit {
				break
			}
			processed++
			if err := s.syncer.SyncOrder(ctx, id); err != nil {
				logger.Error("sync order", "order", id, "error", err)
				continue
			}
			saved++
		}
		// Advance by the number actually returned (self-adjusting to the API's
		// page size) and stop once we've caught up to total or hit the cap.
		offset += len(ids)
		if offset >= total || (s.limit > 0 && processed >= s.limit) {
			break
		}
	}
	logger.Info("lazada order sync done", "processed", processed, "saved", saved)
	return nil
}

// ── small env helpers (test-friendly overrides) ─────────────────────────────

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return fallback
}
