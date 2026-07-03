// Package lazada assembles the Lazada sell-channel. Each entrypoint wires only
// the infrastructure its role needs — the consumer never opens a Kafka
// producer, the scheduler never starts an HTTP server — so one role's missing
// dependency can't stop an unrelated one from starting. Shared setup lives in
// the private helpers below.
package lazada

import (
	"context"

	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	pgAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	redisAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/redis"
	orderUC "github.com/okdev/marketplace-sync/internal/usecase/order"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/catalogscheduler"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/catalogsync"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/client"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/consumer"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/fulfillmentscheduler"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/fulfillmentsync"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/orderscheduler"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/ordersync"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/productcreatescheduler"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/productcreatesync"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/productoffsalescheduler"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/productoffsalesync"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/productscheduler"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/productsync"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/productupdatescheduler"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/productupdatesync"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/sellablestockscheduler"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/sellablestocksync"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/server"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/tokenstore"
)

// NewServer wires the Lazada HTTP server: Postgres (orders), Redis (token store
// used by the webhook handler) and a Kafka producer.
func NewServer(pgCfg pgAdapter.Config, redisCfg redisAdapter.Config, kafkaCfg kafkaAdapter.Config) (*server.Server, error) {
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
	processOrderUC := orderUC.NewProcessWebhook(pgAdapter.NewOrderRepository(db))
	return server.New(processOrderUC, rdb, producer), nil
}

// NewOrderScheduler wires the order-sync scheduler: Postgres (sync bookkeeping)
// and Redis (token store). No Kafka, no HTTP server.
func NewOrderScheduler(cfg client.Config, pgCfg pgAdapter.Config, redisCfg redisAdapter.Config) (*orderscheduler.Scheduler, error) {
	syncer, tokens, err := newSyncer(cfg, pgCfg, redisCfg)
	if err != nil {
		return nil, err
	}
	return orderscheduler.New(syncer, tokens), nil
}

// NewConsumer wires the webhook->order consumer: Postgres (sync bookkeeping),
// Redis (token store) and the Kafka consumer. No Kafka producer, no HTTP server.
func NewConsumer(cfg client.Config, pgCfg pgAdapter.Config, redisCfg redisAdapter.Config, kafkaCfg kafkaAdapter.Config) (*consumer.Consumer, error) {
	syncer, _, err := newSyncer(cfg, pgCfg, redisCfg)
	if err != nil {
		return nil, err
	}
	return consumer.New(kafkaCfg, syncer), nil
}

// NewProductScheduler wires the product price/stock sync scheduler: Postgres
// (product source table) and Redis (token store). No Kafka, no HTTP server.
func NewProductScheduler(cfg client.Config, pgCfg pgAdapter.Config, redisCfg redisAdapter.Config) (*productscheduler.Scheduler, error) {
	db, err := pgAdapter.New(pgCfg)
	if err != nil {
		return nil, err
	}
	rdb, err := redisAdapter.New(redisCfg)
	if err != nil {
		return nil, err
	}
	tokens := tokenstore.New(rdb)
	repo := pgAdapter.NewLazadaProductRepository(db)
	if err := repo.EnsureSchema(context.Background()); err != nil {
		return nil, err
	}
	return productscheduler.New(productsync.New(cfg, tokens, repo), tokens), nil
}

// NewProductCreateScheduler wires the product-create scheduler: Postgres
// (products pending creation) and Redis (token store). No Kafka, no HTTP server.
func NewProductCreateScheduler(cfg client.Config, pgCfg pgAdapter.Config, redisCfg redisAdapter.Config) (*productcreatescheduler.Scheduler, error) {
	db, err := pgAdapter.New(pgCfg)
	if err != nil {
		return nil, err
	}
	rdb, err := redisAdapter.New(redisCfg)
	if err != nil {
		return nil, err
	}
	tokens := tokenstore.New(rdb)
	repo := pgAdapter.NewLazadaProductCreateRepository(db)
	if err := repo.EnsureSchema(context.Background()); err != nil {
		return nil, err
	}
	return productcreatescheduler.New(productcreatesync.New(cfg, tokens, repo), tokens), nil
}

// NewCatalogScheduler wires the catalog reconcile scheduler: Postgres (snapshot)
// and Redis (token store). Inbound poll; no Kafka, no HTTP server.
func NewCatalogScheduler(cfg client.Config, pgCfg pgAdapter.Config, redisCfg redisAdapter.Config) (*catalogscheduler.Scheduler, error) {
	db, err := pgAdapter.New(pgCfg)
	if err != nil {
		return nil, err
	}
	rdb, err := redisAdapter.New(redisCfg)
	if err != nil {
		return nil, err
	}
	tokens := tokenstore.New(rdb)
	repo := pgAdapter.NewLazadaProductCatalogRepository(db)
	if err := repo.EnsureSchema(context.Background()); err != nil {
		return nil, err
	}
	return catalogscheduler.New(catalogsync.New(cfg, tokens, repo), tokens), nil
}

// NewProductUpdateScheduler wires the product-update scheduler: Postgres
// (edits pending push) and Redis (token store). No Kafka, no HTTP server.
func NewProductUpdateScheduler(cfg client.Config, pgCfg pgAdapter.Config, redisCfg redisAdapter.Config) (*productupdatescheduler.Scheduler, error) {
	db, err := pgAdapter.New(pgCfg)
	if err != nil {
		return nil, err
	}
	rdb, err := redisAdapter.New(redisCfg)
	if err != nil {
		return nil, err
	}
	tokens := tokenstore.New(rdb)
	repo := pgAdapter.NewLazadaProductUpdateRepository(db)
	if err := repo.EnsureSchema(context.Background()); err != nil {
		return nil, err
	}
	return productupdatescheduler.New(productupdatesync.New(cfg, tokens, repo), tokens), nil
}

// NewSellableStockScheduler wires the sellable-stock scheduler: Postgres
// (pending stock changes) and Redis (token store). No Kafka, no HTTP server.
func NewSellableStockScheduler(cfg client.Config, pgCfg pgAdapter.Config, redisCfg redisAdapter.Config) (*sellablestockscheduler.Scheduler, error) {
	db, err := pgAdapter.New(pgCfg)
	if err != nil {
		return nil, err
	}
	rdb, err := redisAdapter.New(redisCfg)
	if err != nil {
		return nil, err
	}
	tokens := tokenstore.New(rdb)
	repo := pgAdapter.NewLazadaSellableStockRepository(db)
	if err := repo.EnsureSchema(context.Background()); err != nil {
		return nil, err
	}
	return sellablestockscheduler.New(sellablestocksync.New(cfg, tokens, repo), tokens), nil
}

// NewProductOffsaleScheduler wires the off-sale scheduler: Postgres (pending
// take-down actions) and Redis (token store). No Kafka, no HTTP server.
func NewProductOffsaleScheduler(cfg client.Config, pgCfg pgAdapter.Config, redisCfg redisAdapter.Config) (*productoffsalescheduler.Scheduler, error) {
	db, err := pgAdapter.New(pgCfg)
	if err != nil {
		return nil, err
	}
	rdb, err := redisAdapter.New(redisCfg)
	if err != nil {
		return nil, err
	}
	tokens := tokenstore.New(rdb)
	repo := pgAdapter.NewLazadaProductOffsaleRepository(db)
	if err := repo.EnsureSchema(context.Background()); err != nil {
		return nil, err
	}
	return productoffsalescheduler.New(productoffsalesync.New(cfg, tokens, repo), tokens), nil
}

// NewFulfillmentScheduler wires the fulfillment scheduler: Postgres (orders
// pending fulfillment) and Redis (token store). No Kafka, no HTTP server.
func NewFulfillmentScheduler(cfg client.Config, pgCfg pgAdapter.Config, redisCfg redisAdapter.Config) (*fulfillmentscheduler.Scheduler, error) {
	db, err := pgAdapter.New(pgCfg)
	if err != nil {
		return nil, err
	}
	rdb, err := redisAdapter.New(redisCfg)
	if err != nil {
		return nil, err
	}
	tokens := tokenstore.New(rdb)
	repo := pgAdapter.NewLazadaFulfillmentRepository(db)
	if err := repo.EnsureSchema(context.Background()); err != nil {
		return nil, err
	}
	return fulfillmentscheduler.New(fulfillmentsync.New(cfg, tokens, repo), tokens), nil
}

// newSyncer builds the order syncer shared by the scheduler and the consumer,
// plus the token store they both need.
func newSyncer(cfg client.Config, pgCfg pgAdapter.Config, redisCfg redisAdapter.Config) (*ordersync.Syncer, *tokenstore.Store, error) {
	db, err := pgAdapter.New(pgCfg)
	if err != nil {
		return nil, nil, err
	}
	rdb, err := redisAdapter.New(redisCfg)
	if err != nil {
		return nil, nil, err
	}
	tokens := tokenstore.New(rdb)
	syncRepo := pgAdapter.NewLazadaSyncRepository(db)
	if err := syncRepo.EnsureSchema(context.Background()); err != nil {
		return nil, nil, err
	}
	return ordersync.New(cfg, tokens, syncRepo), tokens, nil
}
