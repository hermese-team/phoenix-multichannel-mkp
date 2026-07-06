// Package shopee assembles the Shopee sell-channel. Each entrypoint wires only
// the infrastructure its role needs.
package shopee

import (
	"context"
	"fmt"
	"log"

	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	pgAdapter    "github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	redisAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/redis"
	orderUC      "github.com/okdev/marketplace-sync/internal/usecase/order"
	productUC    "github.com/okdev/marketplace-sync/internal/usecase/product"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/client"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/consumer"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/scheduler"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/server"
)

// NewServer wires the HTTP server: Postgres (orders), Redis (token store),
// Kafka producer (debug publish), and Shopee API client (OAuth flow).
func NewServer(cfg Config, pgCfg pgAdapter.Config, redisCfg redisAdapter.Config, kafkaCfg kafkaAdapter.Config) (*server.Server, error) {
	log.Printf("[shopee] connecting to postgres DSN: %s", pgCfg.DSN)
	db, err := pgAdapter.New(pgCfg)
	if err != nil {
		return nil, err
	}
	var currentDB string
	_ = db.QueryRow("SELECT current_database()").Scan(&currentDB)
	log.Printf("[shopee] connected to database: %s", currentDB)
	rdb, err := redisAdapter.New(redisCfg)
	if err != nil {
		return nil, err
	}
	producer, err := kafkaAdapter.NewProducer(kafkaCfg)
	if err != nil {
		return nil, err
	}
	tokenRepo := pgAdapter.NewShopeeTokenRepository(db)
	log.Printf("[shopee] creating shopee_tokens table if not exists...")
	if err := tokenRepo.EnsureSchema(context.Background()); err != nil {
		return nil, fmt.Errorf("ensure shopee schema: %w", err)
	}
	log.Printf("[shopee] shopee_tokens table ready")
	shopeeClient := client.New(client.Config{
		PartnerID: cfg.PartnerID,
		AppKey:    cfg.AppKey,
		AppSecret: cfg.AppSecret,
		BaseURL:   cfg.BaseURL,
	})
	processOrderUC := orderUC.NewProcessWebhook(pgAdapter.NewOrderRepository(db))
	return server.New(processOrderUC, shopeeClient, rdb, tokenRepo, producer), nil
}

// NewConsumer wires the product-publish consumer.
func NewConsumer(cfg Config, pgCfg pgAdapter.Config, kafkaCfg kafkaAdapter.Config) (*consumer.ProductConsumer, error) {
	db, err := pgAdapter.New(pgCfg)
	if err != nil {
		return nil, err
	}
	shopeeClient := client.New(client.Config{
		PartnerID: cfg.PartnerID,
		AppKey:    cfg.AppKey,
		AppSecret: cfg.AppSecret,
		BaseURL:   cfg.BaseURL,
	})
	publishUC := productUC.NewPublish(pgAdapter.NewProductRepository(db), shopeeClient)
	return consumer.New(kafkaCfg, publishUC), nil
}

// NewScheduler wires the Shopee scheduler.
func NewScheduler() *scheduler.Scheduler {
	return scheduler.New()
}
