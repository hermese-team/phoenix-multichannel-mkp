package infrastructure

import (
	"git.amaze-x.com/phoenix/channel-adapter-shopee/config"
	"github.com/hermese-team/phoenix-multichannel-mkp/libraries/go/telemetry"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.uber.org/zap"
)

// NewLogger delegates to the shared telemetry library.
// Callers must defer lp.Shutdown(ctx) on service exit.
func NewLogger(cfg config.Config) (*zap.Logger, *sdklog.LoggerProvider, error) {
	return telemetry.NewLogger(telemetry.Options{
		Endpoint:    cfg.Telemetry.Endpoint,
		ServiceName: cfg.Telemetry.ServiceName,
		Env:         cfg.AppEnv,
		LogLevel:    cfg.Logger.LogLevel,
	})
}
