package telemetry

// Options configures the shared OTel tracer and logger providers.
type Options struct {
	Endpoint    string // OTLP gRPC endpoint, e.g. "localhost:4317"
	ServiceName string
	Env         string // "local", "staging", "production"
	LogLevel    string // "debug", "info", "warn", "error"
}
