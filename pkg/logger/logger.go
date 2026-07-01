package logger

import (
	"github.com/ascend/phoenix-multichannel-mkp/config"
	"github.com/hermese-team/phoenix-multichannel-mkp/libraries/go/telemetry"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.uber.org/zap"
)

// New creates a *zap.Logger backed by the shared telemetry library.
// When cfg.Telemetry.Endpoint is set, logs are forwarded to OTLP (SigNoz).
// Callers must defer lp.Shutdown(ctx) on service exit.
func New(cfg *config.Config) (*zap.Logger, *sdklog.LoggerProvider, error) {
	return telemetry.NewLogger(telemetry.Options{
		Endpoint:    cfg.Telemetry.Endpoint,
		ServiceName: cfg.Telemetry.ServiceName,
		Env:         cfg.App.Env,
		LogLevel:    cfg.App.LogLevel,
	})
}
