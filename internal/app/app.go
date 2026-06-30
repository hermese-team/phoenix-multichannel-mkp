package app

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	"github.com/ascend/phoenix-multichannel-mkp/config"
	"github.com/ascend/phoenix-multichannel-mkp/internal/handler"
	"github.com/ascend/phoenix-multichannel-mkp/internal/infrastructure/httpclient"
	"github.com/ascend/phoenix-multichannel-mkp/internal/infrastructure/kafka"
	"github.com/ascend/phoenix-multichannel-mkp/internal/infrastructure/marketplace"
	pginfra "github.com/ascend/phoenix-multichannel-mkp/internal/infrastructure/postgres"
	redisinfra "github.com/ascend/phoenix-multichannel-mkp/internal/infrastructure/redis"
	orderrepo "github.com/ascend/phoenix-multichannel-mkp/internal/repository/order"
	productrepo "github.com/ascend/phoenix-multichannel-mkp/internal/repository/product"
	ordersvc "github.com/ascend/phoenix-multichannel-mkp/internal/service/order"
	productsvc "github.com/ascend/phoenix-multichannel-mkp/internal/service/product"
	"github.com/ascend/phoenix-multichannel-mkp/pkg/logger"
)

// App is the dependency-injection container. It owns every long-lived
// resource (DB pool, cache, broker) and the HTTP server, and is responsible
// for releasing them on shutdown.
type App struct {
	cfg      *config.Config
	pool     *pgxpool.Pool
	redis    *redis.Client
	producer *kafka.Producer
	server   *handler.Server
}

// New builds the full dependency graph from configuration.
func New(ctx context.Context, cfg *config.Config) (*App, error) {
	// ── Infrastructure ──────────────────────────────────────────────────────
	pool, err := pginfra.Connect(ctx, cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("connect postgres: %w", err)
	}
	logger.Info("connected to postgres", "host", cfg.Postgres.Host)

	rdb, err := redisinfra.Connect(ctx, cfg.Redis)
	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("connect redis: %w", err)
	}
	logger.Info("connected to redis", "addr", cfg.Redis.Addr())

	producer, err := kafka.NewProducer(ctx, cfg.Kafka)
	if err != nil {
		pool.Close()
		_ = rdb.Close()
		return nil, fmt.Errorf("connect kafka: %w", err)
	}
	logger.Info("connected to kafka", "brokers", cfg.Kafka.Brokers)

	// ── Outbound marketplace gateways (resty) ───────────────────────────────
	restyClient := httpclient.New(15 * time.Second)
	marketplaces := marketplace.NewRegistry(
		marketplace.NewTikTokGateway(restyClient, cfg.Marketplace.TikTok),
		marketplace.NewLazadaGateway(restyClient, cfg.Marketplace.Lazada),
		marketplace.NewShopeeGateway(restyClient, cfg.Marketplace.Shopee),
	)

	// ── Repositories (inject infra) ─────────────────────────────────────────
	productRepo := productrepo.NewProductRepository(pool, rdb)
	orderRepo := orderrepo.NewOrderRepository(pool)

	// ── Services (inject repos + publisher + gateways) ──────────────────────
	productService := productsvc.NewProductService(productRepo, marketplaces)
	orderService := ordersvc.NewOrderService(orderRepo, productRepo, producer, cfg.Kafka.Topics.OrderCreated)

	// ── Handlers + HTTP server ──────────────────────────────────────────────
	productHandler := handler.NewProductHandler(productService)
	orderHandler := handler.NewOrderHandler(orderService)
	server := handler.New(handler.Config{
		Port:            cfg.App.Port,
		ShutdownTimeout: 10 * time.Second,
	}, productHandler, orderHandler)

	return &App{
		cfg:      cfg,
		pool:     pool,
		redis:    rdb,
		producer: producer,
		server:   server,
	}, nil
}

// Start runs the HTTP server (blocking).
func (a *App) Start() error {
	return a.server.Start()
}

// Shutdown gracefully stops the server and releases all infrastructure.
func (a *App) Shutdown(ctx context.Context) error {
	err := a.server.Shutdown(ctx)

	a.producer.Close()
	if cerr := a.redis.Close(); cerr != nil {
		logger.Error("close redis", "error", cerr)
	}
	a.pool.Close()

	return err
}
