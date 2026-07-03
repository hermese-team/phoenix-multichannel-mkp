// Package catalogscheduler periodically reconciles the Lazada product catalog
// into the local snapshot via the shared catalogsync.Syncer.
package catalogscheduler

import (
	"context"
	"fmt"
	"os"

	"github.com/robfig/cron/v3"

	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/catalogsync"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/tokenstore"
)

type Scheduler struct {
	cron   *cron.Cron
	syncer *catalogsync.Syncer
	tokens *tokenstore.Store

	spec   string
	filter string
}

func New(syncer *catalogsync.Syncer, tokens *tokenstore.Store) *Scheduler {
	return &Scheduler{
		cron:   cron.New(),
		syncer: syncer,
		tokens: tokens,
		spec:   getEnv("LAZADA_CATALOG_SPEC", "@every 1h"),
		filter: getEnv("LAZADA_CATALOG_FILTER", "all"),
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
	logger.Info("lazada catalog scheduler started", "spec", s.spec, "filter", s.filter)

	<-ctx.Done()
	s.cron.Stop()
	logger.Info("lazada catalog scheduler stopped")
	return nil
}

func (s *Scheduler) run(ctx context.Context) {
	if err := s.syncer.Reconcile(ctx, s.filter); err != nil {
		logger.Error("lazada catalog reconcile failed", "error", err)
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
