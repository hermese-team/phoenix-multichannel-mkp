// Package telemetry provides shared OpenTelemetry initialisation helpers for all
// phoenix-multichannel-mkp services.
//
// Usage:
//
//	opts := telemetry.Options{
//	    Endpoint:    cfg.Telemetry.Endpoint,
//	    ServiceName: cfg.Telemetry.ServiceName,
//	    Env:         cfg.App.Env,
//	    LogLevel:    cfg.App.LogLevel,
//	}
//	tp, err := telemetry.NewTracer(opts)
//	log, lp, err := telemetry.NewLogger(opts)
package telemetry

// Options holds the configuration for both tracer and logger setup.
// All fields are optional; an empty Endpoint disables OTLP export (no-op mode).
type Options struct {
	// Endpoint is the OTLP gRPC collector address, e.g. "localhost:4317".
	// When empty, a no-op provider is used — safe for local dev with no SigNoz.
	Endpoint string

	// ServiceName is the OTel service.name resource attribute.
	ServiceName string

	// Env is the deployment environment string, e.g. "dev", "staging", "production".
	// Controls zap encoder selection (development vs production) and OTel resource.
	Env string

	// LogLevel is the minimum zap log level in production mode: "debug", "info", "warn", "error".
	// Ignored when Env is "dev" or "development" (zap.NewDevelopment is used instead).
	LogLevel string
}
