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
	"github.com/okdev/marketplace-sync/sellchannel/shopee/classifier"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/client"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/consumer"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/fulfillmentscheduler"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/fulfillmentsync"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/scheduler"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/server"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/tokenstore"
)

// shopeeRefresher adapts client.Client to the tokenstore.Refresher interface.
type shopeeRefresher struct{ c *client.Client }

func (r *shopeeRefresher) RefreshToken(ctx context.Context, shopID int64, refreshToken string) (string, string, int64, error) {
	resp, err := r.c.RefreshToken(ctx, shopID, refreshToken)
	if err != nil {
		return "", "", 0, err
	}
	if resp.Error != "" {
		return "", "", 0, fmt.Errorf("shopee refresh: %s — %s", resp.Error, resp.Message)
	}
	return resp.AccessToken, resp.RefreshToken, resp.ExpireIn, nil
}

func newTokenStore(repo *pgAdapter.ShopeeTokenRepository, rdb *redisAdapter.Client, c *client.Client) *tokenstore.Store {
	return tokenstore.New(repo, rdb, &shopeeRefresher{c: c})
}

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
	tokens := newTokenStore(tokenRepo, rdb, shopeeClient)
	processOrderUC := orderUC.NewProcessWebhook(pgAdapter.NewOrderRepository(db))
	return server.New(processOrderUC, shopeeClient, rdb, tokens, producer, server.Config{
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

// NewIngestionConsumer wires the order ingestion consumer.
// It reads order.enriched.v1, upserts to the orders table (idempotent ON CONFLICT),
// and publishes order.received.v1 for downstream consumers.
func NewIngestionConsumer(pgCfg pgAdapter.Config, kafkaCfg kafkaAdapter.Config) (*consumer.IngestionConsumer, error) {
	db, err := pgAdapter.New(pgCfg)
	if err != nil {
		return nil, err
	}
	orderRepo := pgAdapter.NewOrderRepository(db)
	if err := orderRepo.EnsureSchema(context.Background()); err != nil {
		return nil, fmt.Errorf("ensure orders schema: %w", err)
	}
	producer, err := kafkaAdapter.NewProducer(kafkaCfg)
	if err != nil {
		return nil, fmt.Errorf("ingestion producer: %w", err)
	}
	return consumer.NewIngestionConsumer(kafkaCfg, orderRepo, producer), nil
}

// NewClassifierConsumer wires the event classifier.
// It reads shopee.order.raw, maps Shopee push codes to canonical EventTypes,
// and publishes IngestEvents to order.ingest.shopee.v1.
func NewClassifierConsumer(kafkaCfg kafkaAdapter.Config) (*classifier.Consumer, error) {
	producer, err := kafkaAdapter.NewProducer(kafkaCfg)
	if err != nil {
		return nil, fmt.Errorf("classifier producer: %w", err)
	}
	return classifier.NewConsumer(kafkaCfg, producer), nil
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
	shopeeClient := client.New(client.Config{
		PartnerID: cfg.PartnerID,
		AppKey:    cfg.AppKey,
		AppSecret: cfg.AppSecret,
		BaseURL:   cfg.BaseURL,
	})
	tokens := newTokenStore(tokenRepo, rdb, shopeeClient)
	fulfillmentRepo := pgAdapter.NewShopeeFulfillmentRepository(db)
	if err := fulfillmentRepo.EnsureSchema(context.Background()); err != nil {
		return nil, fmt.Errorf("ensure shopee fulfillment schema: %w", err)
	}
	return consumer.NewOrderConsumer(kafkaCfg, shopeeClient, tokens, producer, fulfillmentRepo), nil
}

// NewFulfillmentScheduler wires the Shopee fulfillment scheduler.
func NewFulfillmentScheduler(cfg Config, pgCfg pgAdapter.Config, redisCfg redisAdapter.Config, kafkaCfg kafkaAdapter.Config) (*fulfillmentscheduler.Scheduler, error) {
	db, err := pgAdapter.New(pgCfg)
	if err != nil {
		return nil, err
	}
	rdb, err := redisAdapter.New(redisCfg)
	if err != nil {
		return nil, err
	}
	tokenRepo := pgAdapter.NewShopeeTokenRepository(db)
	shopeeClient := client.New(client.Config{
		PartnerID: cfg.PartnerID,
		AppKey:    cfg.AppKey,
		AppSecret: cfg.AppSecret,
		BaseURL:   cfg.BaseURL,
	})
	tokens := newTokenStore(tokenRepo, rdb, shopeeClient)
	fulfillmentRepo := pgAdapter.NewShopeeFulfillmentRepository(db)
	if err := fulfillmentRepo.EnsureSchema(context.Background()); err != nil {
		return nil, fmt.Errorf("ensure shopee fulfillment schema: %w", err)
	}
	syncer := fulfillmentsync.New(shopeeClient, tokens, fulfillmentRepo)
	return fulfillmentscheduler.New(syncer, cfg.FulfillmentSpec, cfg.FulfillmentLimit), nil
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
	shopeeClient := client.New(client.Config{
		PartnerID: cfg.PartnerID,
		AppKey:    cfg.AppKey,
		AppSecret: cfg.AppSecret,
		BaseURL:   cfg.BaseURL,
	})
	tokens := newTokenStore(tokenRepo, rdb, shopeeClient)
	cursorRepo := pgAdapter.NewShopeePollCursorRepository(db)
	if err := cursorRepo.EnsureSchema(context.Background()); err != nil {
		return nil, fmt.Errorf("ensure shopee_poll_cursors schema: %w", err)
	}
	scanRepo := pgAdapter.NewSafetyNetScanRepository(db)
	if err := scanRepo.EnsureSchema(context.Background()); err != nil {
		return nil, fmt.Errorf("ensure safety_net_scan_results schema: %w", err)
	}
	pollJob := scheduler.NewOrderPollJob(shopeeClient, tokens, tokenRepo, cursorRepo, producer, rdb)
	scanJob := scheduler.NewSafetyNetScanJob(shopeeClient, tokens, tokenRepo, scanRepo, producer, rdb)
	return scheduler.New(pollJob, cfg.PollSpec, scanJob, cfg.ScanSpec), nil
}
