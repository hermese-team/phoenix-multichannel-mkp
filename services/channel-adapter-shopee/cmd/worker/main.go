package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.amaze-x.com/phoenix/channel-adapter-shopee/config"
	"git.amaze-x.com/phoenix/channel-adapter-shopee/internal/infrastructure"
	"git.amaze-x.com/phoenix/channel-adapter-shopee/internal/shopee"
	"git.amaze-x.com/phoenix/channel-adapter-shopee/internal/worker"
	"go.uber.org/zap"
)

func main() {
	cfg := config.MustLoad()

	// Logger (with OTel log bridge when telemetry endpoint is set)
	log, lp, err := infrastructure.NewLogger(cfg)
	if err != nil {
		panic("logger init failed: " + err.Error())
	}
	defer log.Sync()
	defer lp.Shutdown(context.Background()) //nolint:errcheck

	log.Info("channel-adapter-shopee worker starting", zap.String("config", config.RedactedJSON(cfg)))

	// Tracer
	tp, err := infrastructure.NewTracer(cfg)
	if err != nil {
		log.Fatal("tracer init failed", zap.Error(err))
	}

	// Redis
	rdb, err := infrastructure.NewRedis(cfg)
	if err != nil {
		log.Fatal("redis init failed", zap.Error(err))
	}

	// Kafka producer
	producer, err := infrastructure.NewKafkaProducer(cfg)
	if err != nil {
		log.Fatal("kafka producer init failed", zap.Error(err))
	}
	defer producer.Close()

	// Shopee signing + token store
	signer := shopee.NewSigner(cfg.Shopee.PartnerID, cfg.Shopee.PartnerKey)
	tokenStore := shopee.NewRedisTokenStore(rdb, cfg.Shopee.BaseURL, signer,
		cfg.Shopee.PartnerID, cfg.Shopee.ShopID)

	// Seed initial tokens from Vault-injected config values.
	if cfg.Shopee.AccessToken != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		if err := tokenStore.Seed(ctx, cfg.Shopee.AccessToken, cfg.Shopee.RefreshToken, 4*time.Hour); err != nil {
			log.Warn("token seed failed — will refresh on first API call", zap.Error(err))
		}
		cancel()
	}

	// Shopee HTTP client + order API
	shopeeClient := shopee.NewClient(
		cfg.Shopee.BaseURL,
		time.Duration(cfg.Shopee.TimeoutSec)*time.Second,
		signer,
		cfg.Shopee.ShopID,
		tokenStore,
	)
	orderAPI := shopee.NewOrderAPI(shopeeClient)

	// Poller
	poller := worker.NewPoller(cfg, orderAPI, producer, rdb, log)

	// Graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		log.Info("shutdown signal received")
		cancel()
	}()

	if err := poller.Run(ctx); err != nil && err != context.Canceled {
		log.Error("poller exited with error", zap.Error(err))
	}

	shutCtx, shutCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutCancel()
	if err := tp.Shutdown(shutCtx); err != nil {
		log.Error("tracer shutdown error", zap.Error(err))
	}

	log.Info("worker stopped")
}
