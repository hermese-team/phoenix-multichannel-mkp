package kafka

import (
	"context"

	"github.com/ascend/phoenix-multichannel-mkp/config"
	"github.com/hermese-team/phoenix-multichannel-mkp/libraries/go/eventing"
)

// Producer adapts the shared eventing.Producer to the service.EventPublisher interface.
// SASL credentials can be added to KafkaConfig when needed; currently local/dev uses no auth.
type Producer struct {
	p *eventing.Producer
}

// NewProducer creates a Kafka producer via the shared eventing library.
// ctx is currently unused by the library (ping timeout is internal) but kept for API consistency.
func NewProducer(_ context.Context, cfg config.KafkaConfig) (*Producer, error) {
	p, err := eventing.NewProducer(eventing.ProducerOptions{
		Brokers: cfg.Brokers,
		// Username/Password: add to KafkaConfig when SASL is required in non-local environments.
	})
	if err != nil {
		return nil, err
	}
	return &Producer{p: p}, nil
}

// Publish delegates to eventing.Producer.Produce and satisfies service.EventPublisher.
func (p *Producer) Publish(ctx context.Context, topic, key string, payload []byte) error {
	return p.p.Produce(ctx, topic, key, payload)
}

// Close flushes pending records and shuts down the underlying client.
func (p *Producer) Close() {
	p.p.Close()
}
