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
	"github.com/okdev/marketplace-sync/sellchannel/lazada/client"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/consumer"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/ordersync"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/scheduler"
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

// NewScheduler wires the order-sync scheduler: Postgres (sync bookkeeping) and
// Redis (token store). No Kafka, no HTTP server.
func NewScheduler(cfg client.Config, pgCfg pgAdapter.Config, redisCfg redisAdapter.Config) (*scheduler.Scheduler, error) {
	syncer, tokens, err := newSyncer(cfg, pgCfg, redisCfg)
	if err != nil {
		return nil, err
	}
	return scheduler.New(syncer, tokens), nil
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
