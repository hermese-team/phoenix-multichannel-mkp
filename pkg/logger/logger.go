package logger

import (
	"context"
	"fmt"

	"github.com/ascend/phoenix-multichannel-mkp/config"
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

// New creates a *zap.Logger. When cfg.Telemetry.Endpoint is set, logs are also
// forwarded to SigNoz/OTLP via otelzap bridge so traces and logs appear together.
// Returns the logger, the log provider (call Shutdown on exit), and any error.
func New(cfg *config.Config) (*zap.Logger, *sdklog.LoggerProvider, error) {
	base, err := buildBase(cfg.App.Env, cfg.App.LogLevel)
	if err != nil {
		return nil, nil, err
	}

	if cfg.Telemetry.Endpoint == "" {
		return base, sdklog.NewLoggerProvider(), nil
	}

	// OTel log provider — shares the same gRPC endpoint as the tracer.
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
			semconv.DeploymentEnvironmentKey.String(cfg.App.Env),
		),
	)

	lp := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(exp)),
		sdklog.WithResource(res),
	)

	// Tee: keep writing to stdout AND forward to OTLP.
	otelCore := otelzap.NewCore(cfg.Telemetry.ServiceName, otelzap.WithLoggerProvider(lp))
	combined := zapcore.NewTee(base.Core(), otelCore)
	log := zap.New(combined, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	return log, lp, nil
}

func buildBase(env, level string) (*zap.Logger, error) {
	if env == "development" || env == "dev" {
		return zap.NewDevelopment()
	}
	lvl, err := zapcore.ParseLevel(level)
	if err != nil {
		lvl = zapcore.InfoLevel
	}
	zapCfg := zap.NewProductionConfig()
	zapCfg.Level = zap.NewAtomicLevelAt(lvl)
	zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapCfg.Build()
}
