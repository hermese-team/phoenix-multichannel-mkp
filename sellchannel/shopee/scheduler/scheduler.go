package scheduler

import (
	"context"
	"log"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron     *cron.Cron
	pollJob  *OrderPollJob      // may be nil
	pollSpec string             // cron spec, e.g. "@every 5m"
	scanJob  *SafetyNetScanJob  // may be nil
	scanSpec string             // cron spec, e.g. "@every 15m"
}

func New(pollJob *OrderPollJob, pollSpec string, scanJob *SafetyNetScanJob, scanSpec string) *Scheduler {
	if pollSpec == "" {
		pollSpec = "@every 5m"
	}
	if scanSpec == "" {
		scanSpec = "@every 15m"
	}
	return &Scheduler{
		cron:     cron.New(),
		pollJob:  pollJob,
		pollSpec: pollSpec,
		scanJob:  scanJob,
		scanSpec: scanSpec,
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
	if s.scanJob != nil {
		s.cron.AddFunc(s.scanSpec, s.scanJob.Run)
		log.Printf("shopee safety-net scan registered (spec=%s)", s.scanSpec)
	}
}
