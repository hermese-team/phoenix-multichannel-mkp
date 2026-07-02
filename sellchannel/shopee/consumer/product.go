package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"

	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	productUC "github.com/okdev/marketplace-sync/internal/usecase/product"
)

const (
	productConsumerGroup = "shopee-product-consumer"
	productPublishTopic  = "shopee.product.publish"
)

type ProductConsumer struct {
	cfg       kafkaAdapter.Config
	publishUC *productUC.PublishUsecase
}

type PublishProductEvent struct {
	ShopID    int64 `json:"shop_id"`
	ProductID int64 `json:"product_id"`
}

func New(cfg kafkaAdapter.Config, publishUC *productUC.PublishUsecase) *ProductConsumer {
	return &ProductConsumer{cfg: cfg, publishUC: publishUC}
}

func (c *ProductConsumer) Start(ctx context.Context) error {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(c.cfg.Brokers...),
		kgo.ConsumerGroup(productConsumerGroup),
		kgo.ConsumeTopics(productPublishTopic),
	)
	if err != nil {
		return fmt.Errorf("create kafka consumer: %w", err)
	}
	defer client.Close()

	log.Println("shopee product consumer started")
	for {
		if ctx.Err() != nil {
			return nil
		}
		fetches := client.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			if ctx.Err() != nil {
				return nil
			}
			for _, e := range errs {
				log.Printf("fetch error: %v", e.Err)
			}
			// Back off so a persistent fetch error (broker down, rebalance
			// loop, …) can't spin this loop at full CPU and flood the broker.
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(time.Second):
			}
			continue
		}
		fetches.EachRecord(func(rec *kgo.Record) {
			var event PublishProductEvent
			if err := json.Unmarshal(rec.Value, &event); err != nil {
				log.Printf("unmarshal event error: %v", err)
				return
			}
			if err := c.publishUC.Execute(ctx, event.ShopID, event.ProductID); err != nil {
				log.Printf("publish product error: %v", err)
			}
		})
	}
}
