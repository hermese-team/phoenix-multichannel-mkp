module github.com/hermese-team/phoenix-multichannel-mkp/libraries/go/telemetry

go 1.26

require (
	go.opentelemetry.io/contrib/bridges/otelzap v0.19.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.61.0
	go.opentelemetry.io/otel v1.44.0
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc v0.20.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.44.0
	go.opentelemetry.io/otel/sdk v1.44.0
	go.opentelemetry.io/otel/sdk/log v0.20.0
	go.uber.org/zap v1.28.0
	google.golang.org/grpc v1.71.0
)
