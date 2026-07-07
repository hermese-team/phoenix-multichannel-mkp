package scheduler

import (
	"context"
	"log"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron     *cron.Cron
	pollJob  *OrderPollJob // may be nil
	pollSpec string        // cron spec, e.g. "@every 5m"
}

func New(pollJob *OrderPollJob, pollSpec string) *Scheduler {
	if pollSpec == "" {
		pollSpec = "@every 5m"
	}
	return &Scheduler{
		cron:     cron.New(),
		pollJob:  pollJob,
		pollSpec: pollSpec,
	}
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
	if s.pollJob != nil {
		s.cron.AddFunc(s.pollSpec, s.pollJob.Run)
		log.Printf("shopee order poll registered (spec=%s)", s.pollSpec)
	}
}
