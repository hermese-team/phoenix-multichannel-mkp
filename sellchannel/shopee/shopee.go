package shopee

import (
	mysqlAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/mysql"
	redisAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/redis"
	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	orderUC "github.com/okdev/marketplace-sync/internal/usecase/order"
	productUC "github.com/okdev/marketplace-sync/internal/usecase/product"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/client"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/consumer"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/scheduler"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/server"
)

type Shopee struct {
	server    *server.Server
	consumer  *consumer.ProductConsumer
	scheduler *scheduler.Scheduler
}

func New(cfg Config, mysqlCfg mysqlAdapter.Config, redisCfg redisAdapter.Config, kafkaCfg kafkaAdapter.Config) (*Shopee, error) {
	db, err := mysqlAdapter.New(mysqlCfg)
	if err != nil {
		return nil, err
	}

	if _, err = redisAdapter.New(redisCfg); err != nil {
		return nil, err
	}

	shopeeClient := client.New(client.Config{
		PartnerID: cfg.PartnerID,
		AppKey:    cfg.AppKey,
		AppSecret: cfg.AppSecret,
		BaseURL:   cfg.BaseURL,
	})

	productRepo := mysqlAdapter.NewProductRepository(db)
	orderRepo := mysqlAdapter.NewOrderRepository(db)

	publishUC := productUC.NewPublish(productRepo, shopeeClient)
	processOrderUC := orderUC.NewProcessWebhook(orderRepo)

	return &Shopee{
		server:    server.New(processOrderUC),
		consumer:  consumer.New(kafkaCfg, publishUC),
		scheduler: scheduler.New(),
	}, nil
}

func (s *Shopee) Server() *server.Server              { return s.server }
func (s *Shopee) Consumer() *consumer.ProductConsumer { return s.consumer }
func (s *Shopee) Scheduler() *scheduler.Scheduler     { return s.scheduler }
