// Package eventing provides a Kafka producer with OpenTelemetry trace
// propagation for all phoenix-multichannel-mkp services.
//
// Usage:
//
//	p, err := eventing.NewProducer(eventing.ProducerOptions{
//	    Brokers:  cfg.Kafka.Brokers,
//	    Username: cfg.Kafka.Username,
//	    Password: cfg.Kafka.Password,
//	})
//	defer p.Close()
//	err = p.Produce(ctx, "raw.order.accepted.v1", orderID, payload)
package eventing

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/scram"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
)

const defaultPingTimeout = 10 * time.Second

// ProducerOptions configures a Kafka producer.
type ProducerOptions struct {
	// Brokers is the list of Kafka bootstrap broker addresses.
	Brokers []string

	// Username and Password enable SASL SCRAM-SHA-512 authentication.
	// Leave empty for unauthenticated (local dev) connections.
	Username string
	Password string
}

// Producer wraps a franz-go client with OTel trace propagation.
// Produces with AllISRAcks + Zstd compression for production durability.
type Producer struct {
	client *kgo.Client
}

// NewProducer creates and pings a Kafka producer.
// Returns an error if the brokers are unreachable within 10 seconds.
func NewProducer(opts ProducerOptions) (*Producer, error) {
	kopts := []kgo.Opt{
		kgo.SeedBrokers(opts.Brokers...),
		kgo.RequiredAcks(kgo.AllISRAcks()),
		kgo.ProducerBatchCompression(kgo.ZstdCompression()),
	}
	if opts.Username != "" {
		kopts = append(kopts,
			kgo.DialTLSConfig(&tls.Config{MinVersion: tls.VersionTLS12}),
			kgo.SASL(scram.Auth{
				User: opts.Username,
				Pass: opts.Password,
			}.AsSha512Mechanism()),
		)
	}

	client, err := kgo.NewClient(kopts...)
	if err != nil {
		return nil, fmt.Errorf("eventing: create kafka client: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), defaultPingTimeout)
	defer cancel()
	if err := client.Ping(pingCtx); err != nil {
		client.Close()
		return nil, fmt.Errorf("eventing: kafka brokers unreachable: %w", err)
	}
	return &Producer{client: client}, nil
}

// Produce publishes a single record synchronously and waits for quorum acks.
// OTel trace context is propagated via Kafka record headers so that consumers
// can resume the same trace.
func (p *Producer) Produce(ctx context.Context, topic, key string, value []byte) error {
	ctx, span := otel.Tracer("eventing.producer").Start(ctx, "kafka.produce")
	defer span.End()
	span.SetAttributes(
		attribute.String("messaging.system", "kafka"),
		attribute.String("messaging.destination", topic),
		attribute.String("messaging.kafka.message_key", key),
	)

	// Inject trace context into record headers.
	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	headers := make([]kgo.RecordHeader, 0, len(carrier))
	for k, v := range carrier {
		headers = append(headers, kgo.RecordHeader{Key: k, Value: []byte(v)})
	}

	rec := &kgo.Record{
		Topic:   topic,
		Key:     []byte(key),
		Value:   value,
		Headers: headers,
	}
	if err := p.client.ProduceSync(ctx, rec).FirstErr(); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return fmt.Errorf("eventing: produce to %s: %w", topic, err)
	}
	return nil
}

// Close flushes pending records and shuts down the client.
func (p *Producer) Close() {
	p.client.Close()
}
