// Package sellablestockscheduler periodically pushes pending warehouse
// sellable-quantity changes to Lazada via the shared sellablestocksync.Syncer.
package sellablestockscheduler

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/robfig/cron/v3"

	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/sellablestocksync"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/tokenstore"
)

type Scheduler struct {
	cron   *cron.Cron
	syncer *sellablestocksync.Syncer
	tokens *tokenstore.Store

	spec  string
	limit int
}

func New(syncer *sellablestocksync.Syncer, tokens *tokenstore.Store) *Scheduler {
	return &Scheduler{
		cron:   cron.New(),
		syncer: syncer,
		tokens: tokens,
		spec:   getEnv("LAZADA_SELLABLE_STOCK_SPEC", "@every 5m"),
		limit:  getEnvInt("LAZADA_SELLABLE_STOCK_LIMIT", 20),
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
	s.run(ctx)

	if _, err := s.cron.AddFunc(s.spec, func() { s.run(context.Background()) }); err != nil {
		return err
	}
	s.cron.Start()
	logger.Info("lazada sellable-stock scheduler started", "spec", s.spec, "limit", s.limit)

	<-ctx.Done()
	s.cron.Stop()
	logger.Info("lazada sellable-stock scheduler stopped")
	return nil
}

func (s *Scheduler) run(ctx context.Context) {
	if err := s.syncer.SyncPending(ctx, s.limit); err != nil {
		logger.Error("lazada sellable-stock failed", "error", err)
	}
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
