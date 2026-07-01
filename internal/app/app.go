package app

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"

	"github.com/ascend/phoenix-multichannel-mkp/config"
	"github.com/ascend/phoenix-multichannel-mkp/internal/handler"
	"github.com/ascend/phoenix-multichannel-mkp/internal/infrastructure"
	"github.com/ascend/phoenix-multichannel-mkp/internal/infrastructure/httpclient"
	"github.com/ascend/phoenix-multichannel-mkp/internal/infrastructure/kafka"
	"github.com/ascend/phoenix-multichannel-mkp/internal/infrastructure/marketplace"
	pginfra "github.com/ascend/phoenix-multichannel-mkp/internal/infrastructure/postgres"
	redisinfra "github.com/ascend/phoenix-multichannel-mkp/internal/infrastructure/redis"
	orderrepo "github.com/ascend/phoenix-multichannel-mkp/internal/repository/order"
	productrepo "github.com/ascend/phoenix-multichannel-mkp/internal/repository/product"
	ordersvc "github.com/ascend/phoenix-multichannel-mkp/internal/service/order"
	productsvc "github.com/ascend/phoenix-multichannel-mkp/internal/service/product"
)

// App is the dependency-injection container. It owns every long-lived
// resource and is responsible for releasing them on shutdown.
type App struct {
	cfg      *config.Config
	log      *zap.Logger
	pool     *pgxpool.Pool
	redis    *redis.Client
	producer *kafka.Producer
	server   *handler.Server
	tp       *sdktrace.TracerProvider
}

// New builds the full dependency graph from configuration.
func New(ctx context.Context, cfg *config.Config, log *zap.Logger) (*App, error) {
	// ── Tracer ──────────────────────────────────────────────────────────────
	tp, err := infrastructure.NewTracer(cfg)
	if err != nil {
		return nil, fmt.Errorf("init tracer: %w", err)
	}

	// ── Postgres ─────────────────────────────────────────────────────────────
	pool, err := pginfra.Connect(ctx, cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("connect postgres: %w", err)
	}
	log.Info("connected to postgres", zap.String("host", cfg.Postgres.Host))

	// ── Migrations ───────────────────────────────────────────────────────────
	if err := pginfra.RunMigrations(cfg.Postgres, "file://migrations"); err != nil {
		pool.Close()
		return nil, fmt.Errorf("run migrations: %w", err)
	}
	log.Info("migrations applied")

	// ── Redis ────────────────────────────────────────────────────────────────
	rdb, err := redisinfra.Connect(ctx, cfg.Redis)
	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("connect redis: %w", err)
	}
	log.Info("connected to redis", zap.String("addr", cfg.Redis.Addr()))

	// ── Kafka ────────────────────────────────────────────────────────────────
	producer, err := kafka.NewProducer(ctx, cfg.Kafka)
	if err != nil {
		pool.Close()
		_ = rdb.Close()
		return nil, fmt.Errorf("connect kafka: %w", err)
	}
	log.Info("connected to kafka", zap.Strings("brokers", cfg.Kafka.Brokers))

	// ── Outbound marketplace gateways ────────────────────────────────────────
	restyClient := httpclient.New(15 * time.Second)
	marketplaces := marketplace.NewRegistry(
		marketplace.NewTikTokGateway(restyClient, cfg.Marketplace.TikTok),
		marketplace.NewLazadaGateway(restyClient, cfg.Marketplace.Lazada),
		marketplace.NewShopeeGateway(restyClient, cfg.Marketplace.Shopee),
	)

	// ── Repositories ─────────────────────────────────────────────────────────
	productRepo := productrepo.NewProductRepository(pool, rdb)
	orderRepo := orderrepo.NewOrderRepository(pool)

	// ── Services ─────────────────────────────────────────────────────────────
	productService := productsvc.NewProductService(productRepo, marketplaces)
	orderService := ordersvc.NewOrderService(orderRepo, productRepo, producer, cfg.Kafka.Topics.OrderCreated)

	// ── HTTP server ──────────────────────────────────────────────────────────
	server := handler.New(
		handler.Config{
			Port:            cfg.App.Port,
			ServiceName:     cfg.Telemetry.ServiceName,
			ShutdownTimeout: 10 * time.Second,
		},
		log,
		handler.NewProductHandler(productService),
		handler.NewOrderHandler(orderService),
	)

	return &App{
		cfg:      cfg,
		log:      log,
		pool:     pool,
		redis:    rdb,
		producer: producer,
		server:   server,
		tp:       tp,
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
		a.log.Error("close redis", zap.Error(cerr))
	}
	a.pool.Close()
	if cerr := a.tp.Shutdown(ctx); cerr != nil {
		a.log.Error("shutdown tracer", zap.Error(cerr))
	}
	return err
}
