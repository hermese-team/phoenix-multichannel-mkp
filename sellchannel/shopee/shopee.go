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
	"github.com/okdev/marketplace-sync/sellchannel/shopee/tokenstore"
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
	return server.New(processOrderUC, shopeeClient, rdb, tokenRepo, producer, server.Config{
		PartnerID:     cfg.PartnerID,
		AppSecret:     cfg.AppSecret,
		WebhookVerify: cfg.WebhookVerify,
	}), nil
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

// NewOrderConsumer wires the order-enrichment consumer.
// It reads raw webhook events from shopee.order.raw,
// fetches full order detail from Shopee API, and publishes to shopee.order.detail.
func NewOrderConsumer(cfg Config, pgCfg pgAdapter.Config, redisCfg redisAdapter.Config, kafkaCfg kafkaAdapter.Config) (*consumer.OrderConsumer, error) {
	db, err := pgAdapter.New(pgCfg)
	if err != nil {
		return nil, err
	}
	rdb, err := redisAdapter.New(redisCfg)
	if err != nil {
		return nil, err
	}
	producer, err := kafkaAdapter.NewProducer(kafkaCfg)
	if err != nil {
		return nil, err
	}
	tokenRepo := pgAdapter.NewShopeeTokenRepository(db)
	tokens := tokenstore.New(tokenRepo, rdb)
	shopeeClient := client.New(client.Config{
		PartnerID: cfg.PartnerID,
		AppKey:    cfg.AppKey,
		AppSecret: cfg.AppSecret,
		BaseURL:   cfg.BaseURL,
	})
	return consumer.NewOrderConsumer(kafkaCfg, shopeeClient, tokens, producer), nil
}

// NewScheduler wires the Shopee fallback-polling scheduler.
// It polls Shopee order list every pollSpec interval to catch orders missed by webhook.
func NewScheduler(cfg Config, pgCfg pgAdapter.Config, redisCfg redisAdapter.Config, kafkaCfg kafkaAdapter.Config) (*scheduler.Scheduler, error) {
	db, err := pgAdapter.New(pgCfg)
	if err != nil {
		return nil, err
	}
	rdb, err := redisAdapter.New(redisCfg)
	if err != nil {
		return nil, err
	}
	producer, err := kafkaAdapter.NewProducer(kafkaCfg)
	if err != nil {
		return nil, err
	}
	tokenRepo := pgAdapter.NewShopeeTokenRepository(db)
	tokens := tokenstore.New(tokenRepo, rdb)
	shopeeClient := client.New(client.Config{
		PartnerID: cfg.PartnerID,
		AppKey:    cfg.AppKey,
		AppSecret: cfg.AppSecret,
		BaseURL:   cfg.BaseURL,
	})
	pollJob := scheduler.NewOrderPollJob(shopeeClient, tokens, tokenRepo, producer, rdb, cfg.PollWindow)
	return scheduler.New(pollJob, cfg.PollSpec), nil
}
