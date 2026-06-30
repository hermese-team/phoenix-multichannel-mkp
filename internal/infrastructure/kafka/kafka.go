package kafka

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/ascend/phoenix-multichannel-mkp/config"
)

// Producer wraps a franz-go client and implements service.EventPublisher.
type Producer struct {
	client *kgo.Client
}

// NewProducer creates a Kafka producer client and pings the cluster.
func NewProducer(ctx context.Context, cfg config.KafkaConfig) (*Producer, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(cfg.Brokers...),
		kgo.AllowAutoTopicCreation(),
	)
	if err != nil {
		return nil, fmt.Errorf("create kafka client: %w", err)
	}
	if err := client.Ping(ctx); err != nil {
		client.Close()
		return nil, fmt.Errorf("ping kafka: %w", err)
	}
	return &Producer{client: client}, nil
}

// Publish synchronously produces a single record to the given topic.
func (p *Producer) Publish(ctx context.Context, topic, key string, payload []byte) error {
	rec := &kgo.Record{
		Topic: topic,
		Key:   []byte(key),
		Value: payload,
	}
	if err := p.client.ProduceSync(ctx, rec).FirstErr(); err != nil {
		return fmt.Errorf("produce to %s: %w", topic, err)
	}
	return nil
}

// Close flushes pending records and closes the client.
func (p *Producer) Close() {
	p.client.Close()
}
