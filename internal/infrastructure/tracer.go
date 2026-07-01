package infrastructure

import (
	"context"
	"fmt"

	"git.amaze-x.com/phoenix/channel-adapter-shopee/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewTracer initialises an OTLP gRPC tracer and registers it as the global provider.
// Returns a shutdown function that must be called on service stop.
func NewTracer(cfg config.Config) (*sdktrace.TracerProvider, error) {
	res, err := resource.New(context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(cfg.Telemetry.ServiceName),
			semconv.DeploymentEnvironmentKey.String(cfg.AppEnv),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("creating OTel resource: %w", err)
	}

	var tp *sdktrace.TracerProvider

	if cfg.Telemetry.Endpoint == "" {
		// No endpoint configured — use a no-op provider (useful in local dev)
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
			return nil, fmt.Errorf("creating OTLP exporter: %w", err)
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
