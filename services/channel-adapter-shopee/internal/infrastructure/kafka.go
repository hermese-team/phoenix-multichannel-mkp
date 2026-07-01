package infrastructure

import (
	"context"

	"git.amaze-x.com/phoenix/channel-adapter-shopee/config"
	"github.com/hermese-team/phoenix-multichannel-mkp/libraries/go/eventing"
)

// KafkaProducer delegates to the shared eventing library.
// OTel trace propagation and SASL/TLS are handled inside the library.
type KafkaProducer struct {
	p *eventing.Producer
}

// NewKafkaProducer creates a Kafka producer via the shared eventing library.
func NewKafkaProducer(cfg config.Config) (*KafkaProducer, error) {
	p, err := eventing.NewProducer(eventing.ProducerOptions{
		Brokers:  cfg.Kafka.Brokers,
		Username: cfg.Kafka.Username,
		Password: cfg.Kafka.Password,
	})
	if err != nil {
		return nil, err
	}
	return &KafkaProducer{p: p}, nil
}

// Produce publishes a single record with OTel trace context in headers.
func (p *KafkaProducer) Produce(ctx context.Context, topic, key string, value []byte) error {
	return p.p.Produce(ctx, topic, key, value)
}

// Close flushes pending records and shuts down the underlying client.
func (p *KafkaProducer) Close() {
	p.p.Close()
}
