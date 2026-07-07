// Package fulfillmentscheduler periodically processes pending Shopee fulfillments.
package fulfillmentscheduler

import (
	"context"
	"log"

	"github.com/robfig/cron/v3"

	"github.com/okdev/marketplace-sync/sellchannel/shopee/fulfillmentsync"
)

type Scheduler struct {
	cron   *cron.Cron
	syncer *fulfillmentsync.Syncer
	spec   string
	limit  int
}

func New(syncer *fulfillmentsync.Syncer, spec string, limit int) *Scheduler {
	if spec == "" {
		spec = "@every 5m"
	}
	if limit <= 0 {
		limit = 20
	}
	return &Scheduler{
		cron:   cron.New(),
		syncer: syncer,
		spec:   spec,
		limit:  limit,
	}
}

func (s *Scheduler) Start(ctx context.Context) {
	// Run once immediately on startup, then on schedule.
	s.run()

	s.cron.AddFunc(s.spec, s.run)
	s.cron.Start()
	log.Printf("shopee fulfillment scheduler started (spec=%s limit=%d)", s.spec, s.limit)

	<-ctx.Done()
	s.cron.Stop()
	log.Println("shopee fulfillment scheduler stopped")
}

func (s *Scheduler) run() {
	if err := s.syncer.SyncPending(context.Background(), s.limit); err != nil {
		log.Printf("[fulfillment-scheduler] sync error: %v", err)
	}
}
