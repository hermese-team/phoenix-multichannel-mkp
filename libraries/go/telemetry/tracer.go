package telemetry

import (
	"context"
	"fmt"

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
// When opts.Endpoint is empty a no-op provider is returned (safe for local dev).
// Callers must defer tp.Shutdown(ctx) on service exit.
func NewTracer(opts Options) (*sdktrace.TracerProvider, error) {
	res, err := newResource(opts)
	if err != nil {
		return nil, err
	}

	var tp *sdktrace.TracerProvider

	if opts.Endpoint == "" {
		tp = sdktrace.NewTracerProvider(sdktrace.WithResource(res))
	} else {
		conn, err := grpc.NewClient(opts.Endpoint,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return nil, fmt.Errorf("telemetry: connecting to OTel collector: %w", err)
		}
		exp, err := otlptracegrpc.New(context.Background(), otlptracegrpc.WithGRPCConn(conn))
		if err != nil {
			return nil, fmt.Errorf("telemetry: creating OTLP trace exporter: %w", err)
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

func newResource(opts Options) (*resource.Resource, error) {
	res, err := resource.New(context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(opts.ServiceName),
			semconv.DeploymentEnvironmentKey.String(opts.Env),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("telemetry: creating OTel resource: %w", err)
	}
	return res, nil
}
