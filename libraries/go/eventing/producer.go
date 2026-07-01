package eventing

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/scram"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// ProducerOptions configures the Kafka producer.
type ProducerOptions struct {
	Brokers  []string
	Username string // SASL SCRAM-SHA-512; leave empty for unauthenticated
	Password string
}

// Producer wraps a franz-go client with OTel trace-context propagation.
type Producer struct {
	client *kgo.Client
}

// NewProducer creates a Kafka producer with optional SASL auth.
func NewProducer(opts ProducerOptions) (*Producer, error) {
	kopts := []kgo.Opt{
		kgo.SeedBrokers(opts.Brokers...),
	}
	if opts.Username != "" {
		kopts = append(kopts, kgo.SASL(scram.Auth{
			User: opts.Username,
			Pass: opts.Password,
		}.AsSha512Mechanism()))
	}
	client, err := kgo.NewClient(kopts...)
	if err != nil {
		return nil, err
	}
	return &Producer{client: client}, nil
}

// Produce publishes one record synchronously. OTel trace context is injected
// into Kafka record headers so consumers can continue the trace.
func (p *Producer) Produce(ctx context.Context, topic, key string, value []byte) error {
	carrier := make(map[string]string)
	otel.GetTextMapPropagator().Inject(ctx, propagation.MapCarrier(carrier))

	headers := make([]kgo.RecordHeader, 0, len(carrier))
	for k, v := range carrier {
		headers = append(headers, kgo.RecordHeader{Key: k, Value: []byte(v)})
	}

	record := &kgo.Record{
		Topic:   topic,
		Key:     []byte(key),
		Value:   value,
		Headers: headers,
	}
	return p.client.ProduceSync(ctx, record).FirstErr()
}

// Close flushes pending records and shuts down the underlying Kafka client.
func (p *Producer) Close() {
	p.client.Close()
}
