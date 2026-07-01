package telemetry

import (
	"context"

	"go.opentelemetry.io/contrib/bridges/otelzap"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewLogger creates a *zap.Logger that tees logs to an OTLP log exporter.
// Callers must defer lp.Shutdown(ctx) on service exit.
func NewLogger(opts Options) (*zap.Logger, *sdklog.LoggerProvider, error) {
	ctx := context.Background()

	conn, err := grpc.NewClient(
		opts.Endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, err
	}

	exporter, err := otlploggrpc.New(ctx, otlploggrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, nil, err
	}

	res, err := newResource(opts)
	if err != nil {
		return nil, nil, err
	}

	lp := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(exporter)),
		sdklog.WithResource(res),
	)

	base, err := buildBase(opts.Env, opts.LogLevel)
	if err != nil {
		return nil, nil, err
	}

	logger := base.WithOptions(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewTee(core, otelzap.NewCore(opts.ServiceName, otelzap.WithLoggerProvider(lp)))
	}))

	return logger, lp, nil
}

func buildBase(env, level string) (*zap.Logger, error) {
	var lvl zapcore.Level
	if err := lvl.UnmarshalText([]byte(level)); err != nil {
		lvl = zapcore.InfoLevel
	}
	if env == "production" || env == "prod" {
		cfg := zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(lvl)
		return cfg.Build()
	}
	cfg := zap.NewDevelopmentConfig()
	cfg.Level = zap.NewAtomicLevelAt(lvl)
	return cfg.Build()
}
