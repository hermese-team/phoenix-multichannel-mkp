package infrastructure

import (
	"git.amaze-x.com/phoenix/channel-adapter-shopee/config"
	"github.com/hermese-team/phoenix-multichannel-mkp/libraries/go/telemetry"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// NewTracer delegates to the shared telemetry library.
// Callers must defer tp.Shutdown(ctx) on service exit.
func NewTracer(cfg config.Config) (*sdktrace.TracerProvider, error) {
	return telemetry.NewTracer(telemetry.Options{
		Endpoint:    cfg.Telemetry.Endpoint,
		ServiceName: cfg.Telemetry.ServiceName,
		Env:         cfg.AppEnv,
	})
}
