package telemetry

import (
	"context"
	"fmt"

	"go.opentelemetry.io/contrib/bridges/otelzap"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewLogger creates a *zap.Logger. When opts.Endpoint is set, logs are also
// forwarded to the OTLP collector via the otelzap bridge so that traces and
// logs are correlated in SigNoz.
//
// Callers must defer lp.Shutdown(ctx) on service exit.
func NewLogger(opts Options) (*zap.Logger, *sdklog.LoggerProvider, error) {
	base, err := buildBase(opts.Env, opts.LogLevel)
	if err != nil {
		return nil, nil, err
	}

	if opts.Endpoint == "" {
		return base, sdklog.NewLoggerProvider(), nil
	}

	conn, err := grpc.NewClient(opts.Endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("telemetry: connecting to OTel log collector: %w", err)
	}

	exp, err := otlploggrpc.New(context.Background(), otlploggrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, nil, fmt.Errorf("telemetry: creating OTLP log exporter: %w", err)
	}

	res, err := newResource(opts)
	if err != nil {
		return nil, nil, err
	}

	lp := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(exp)),
		sdklog.WithResource(res),
	)

	// Tee: keep writing to stdout AND forward to OTLP.
	otelCore := otelzap.NewCore(opts.ServiceName, otelzap.WithLoggerProvider(lp))
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
