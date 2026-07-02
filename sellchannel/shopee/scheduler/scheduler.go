package scheduler

import (
	"context"
	"log"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron *cron.Cron
}

func New() *Scheduler {
	return &Scheduler{cron: cron.New()}
}

func (s *Scheduler) Register(spec string, job func()) error {
	_, err := s.cron.AddFunc(spec, job)
	return err
}

func (s *Scheduler) Start(ctx context.Context) {
	s.registerDefaults()
	s.cron.Start()
	log.Println("shopee scheduler started")
	<-ctx.Done()
	s.cron.Stop()
	log.Println("shopee scheduler stopped")
}

func (s *Scheduler) registerDefaults() {
	s.cron.AddFunc("@every 1h", func() {
		log.Println("refreshing shopee tokens...")
		// TODO: inject and call auth.RefreshTokenUsecase
	})

	s.cron.AddFunc("@every 5m", func() {
		log.Println("retrying failed sync jobs...")
		// TODO: inject and call scheduler.RetryFailedJobUsecase
	})
}
