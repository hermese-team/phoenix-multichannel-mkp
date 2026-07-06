// Package shopee assembles the Shopee sell-channel. Each entrypoint wires only
// the infrastructure its role needs, so the consumer never opens an HTTP
// server and the server never opens a Kafka consumer.
package shopee

import (
	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	pgAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	orderUC "github.com/okdev/marketplace-sync/internal/usecase/order"
	productUC "github.com/okdev/marketplace-sync/internal/usecase/product"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/client"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/consumer"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/scheduler"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/server"
)

// NewServer wires the Shopee HTTP server: Postgres (orders) only.
func NewServer(pgCfg pgAdapter.Config) (*server.Server, error) {
	db, err := pgAdapter.New(pgCfg)
	if err != nil {
		return nil, err
	}
	processOrderUC := orderUC.NewProcessWebhook(pgAdapter.NewOrderRepository(db))
	return server.New(processOrderUC), nil
}

// NewConsumer wires the product-publish consumer: Postgres (products),
// the Shopee client and the Kafka consumer.
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
