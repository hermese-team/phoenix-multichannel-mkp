package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	productUC "github.com/okdev/marketplace-sync/internal/usecase/product"
)

type ProductConsumer struct {
	reader    *kafka.Reader
	publishUC *productUC.PublishUsecase
}

type PublishProductEvent struct {
	ShopID    int64 `json:"shop_id"`
	ProductID int64 `json:"product_id"`
}

func New(cfg kafkaAdapter.Config, publishUC *productUC.PublishUsecase) *ProductConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: cfg.Brokers,
		GroupID: "shopee-product-consumer",
		Topic:   "shopee.product.publish",
	})
	return &ProductConsumer{reader: reader, publishUC: publishUC}
}

func (c *ProductConsumer) Start(ctx context.Context) error {
	log.Println("shopee product consumer started")
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return nil
			}
			log.Printf("read message error: %v", err)
			continue
		}

		var event PublishProductEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("unmarshal event error: %v", err)
			continue
		}

		if err := c.publishUC.Execute(ctx, event.ShopID, event.ProductID); err != nil {
			log.Printf("publish product error: %v", err)
		}
	}
}

func (c *ProductConsumer) Close() error {
	return c.reader.Close()
}
