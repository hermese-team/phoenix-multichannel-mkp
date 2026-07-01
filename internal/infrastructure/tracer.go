package infrastructure

import (
	"context"
	"fmt"

	"github.com/ascend/phoenix-multichannel-mkp/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewTracer initialises an OTLP gRPC trace provider and registers it as the global.
// Returns a shutdown function that must be called on service stop.
// When cfg.Telemetry.Endpoint is empty a no-op provider is used (safe for local dev).
func NewTracer(cfg *config.Config) (*sdktrace.TracerProvider, error) {
	res, err := resource.New(context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(cfg.Telemetry.ServiceName),
			semconv.DeploymentEnvironmentKey.String(cfg.App.Env),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("creating OTel resource: %w", err)
	}

	var tp *sdktrace.TracerProvider

	if cfg.Telemetry.Endpoint == "" {
		tp = sdktrace.NewTracerProvider(sdktrace.WithResource(res))
	} else {
		conn, err := grpc.NewClient(cfg.Telemetry.Endpoint,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return nil, fmt.Errorf("connecting to OTel collector: %w", err)
		}
		exp, err := otlptracegrpc.New(context.Background(), otlptracegrpc.WithGRPCConn(conn))
		if err != nil {
			return nil, fmt.Errorf("creating OTLP trace exporter: %w", err)
		}
		tp = sdktrace.NewTracerProvider(
			sdktrace.WithBatcher(exp),
			sdktrace.WithResource(res),
		)
	}

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))
	return tp, nil
}
