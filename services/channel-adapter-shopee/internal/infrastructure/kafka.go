package infrastructure

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"git.amaze-x.com/phoenix/channel-adapter-shopee/config"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/scram"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
)

const kafkaPingTimeout = 10 * time.Second

func kafkaOpts(cfg config.KafkaConfig) []kgo.Opt {
	opts := []kgo.Opt{kgo.SeedBrokers(cfg.Brokers...)}
	if cfg.Username != "" {
		opts = append(opts,
			kgo.DialTLSConfig(&tls.Config{}),
			kgo.SASL(scram.Auth{
				User: cfg.Username,
				Pass: cfg.Password,
			}.AsSha512Mechanism()),
		)
	}
	return opts
}

// KafkaProducer wraps franz-go with OTel trace propagation.
// Produces with acks=all via ProduceSync for durability guarantee.
type KafkaProducer struct {
	client *kgo.Client
}

func NewKafkaProducer(cfg config.Config) (*KafkaProducer, error) {
	opts := append(kafkaOpts(cfg.Kafka),
		kgo.RequiredAcks(kgo.AllISRAcks()),
		kgo.ProducerBatchCompression(kgo.ZstdCompression()),
	)
	client, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, err
	}
	pingCtx, cancel := context.WithTimeout(context.Background(), kafkaPingTimeout)
	defer cancel()
	if err := client.Ping(pingCtx); err != nil {
		client.Close()
		return nil, fmt.Errorf("kafka brokers unreachable: %w", err)
	}
	return &KafkaProducer{client: client}, nil
}

// Produce publishes a single record and waits for quorum acknowledgement.
// OTel trace context is propagated via Kafka record headers.
func (p *KafkaProducer) Produce(ctx context.Context, topic, key string, value []byte) error {
	ctx, span := otel.Tracer("kafka.producer").Start(ctx, "kafka.produce")
	defer span.End()
	span.SetAttributes(
		attribute.String("messaging.system", "kafka"),
		attribute.String("messaging.destination", topic),
		attribute.String("messaging.kafka.message_key", key),
	)

	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)

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
	if err := p.client.ProduceSync(ctx, record).FirstErr(); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return fmt.Errorf("kafka produce to %s: %w", topic, err)
	}
	return nil
}

func (p *KafkaProducer) Close() {
	p.client.Close()
}
