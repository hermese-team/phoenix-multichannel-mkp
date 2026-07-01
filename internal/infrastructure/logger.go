package infrastructure

import (
	"context"
	"fmt"

	"git.amaze-x.com/phoenix/channel-adapter-shopee/config"
	"go.opentelemetry.io/contrib/bridges/otelzap"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewLogger creates a zap.Logger that also forwards logs to the OTLP endpoint
// (when cfg.Telemetry.Endpoint is set), so logs appear in SigNoz alongside traces.
// Returns the logger, the log provider (for shutdown), and any error.
// If no endpoint is configured the returned provider is a no-op.
func NewLogger(cfg config.Config) (*zap.Logger, *sdklog.LoggerProvider, error) {
	// Base zap logger
	var baseLogger *zap.Logger
	if cfg.AppEnv == "dev" {
		baseLogger, _ = zap.NewDevelopment()
	} else {
		baseLogger, _ = zap.NewProduction()
	}

	if cfg.Telemetry.Endpoint == "" {
		return baseLogger, sdklog.NewLoggerProvider(), nil
	}

	// OTel log provider — reuse same gRPC conn as the tracer
	conn, err := grpc.NewClient(cfg.Telemetry.Endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("connecting to OTel log collector: %w", err)
	}

	exp, err := otlploggrpc.New(context.Background(), otlploggrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, nil, fmt.Errorf("creating OTLP log exporter: %w", err)
	}

	res, _ := resource.New(context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(cfg.Telemetry.ServiceName),
			semconv.DeploymentEnvironmentKey.String(cfg.AppEnv),
		),
	)

	lp := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(exp)),
		sdklog.WithResource(res),
	)

	// Bridge: otelzap.NewCore forwards zap records to the OTel log provider
	otelCore := otelzap.NewCore(cfg.Telemetry.ServiceName, otelzap.WithLoggerProvider(lp))

	// Tee: keep writing to the original zap output AND forward to SigNoz
	combined := zapcore.NewTee(baseLogger.Core(), otelCore)
	logger := zap.New(combined, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	return logger, lp, nil
}
