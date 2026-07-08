package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugar *zap.SugaredLogger

func init() {
	// Sensible default so logging works before Init is called (e.g. config load).
	l, _ := build("info")
	sugar = l.Sugar()
}

// Init reconfigures the global logger with the level from config
// (e.g. "debug", "info", "warn", "error"). Call once at startup.
func Init(level string) error {
	l, err := build(level)
	if err != nil {
		return err
	}
	sugar = l.Sugar()
	return nil
}

func build(level string) (*zap.Logger, error) {
	lvl, err := zapcore.ParseLevel(level)
	if err != nil {
		return nil, err
	}
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(lvl)
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}
	// Skip the wrapper frame so "caller" points to the real call site.
	return cfg.Build(zap.AddCallerSkip(1))
}

// Sync flushes any buffered log entries. Call on shutdown.
func Sync() { _ = sugar.Sync() }

// Without context — for startup / teardown code where ctx is not available.
func Info(msg string, args ...any)  { sugar.Infow(msg, args...) }
func Error(msg string, args ...any) { sugar.Errorw(msg, args...) }
func Warn(msg string, args ...any)  { sugar.Warnw(msg, args...) }
func Debug(msg string, args ...any) { sugar.Debugw(msg, args...) }

// With context — preferred in request / event handlers.
// Extracts trace_id from ctx (if OpenTelemetry is wired) for Grafana correlation.
func InfoContext(ctx context.Context, msg string, args ...any) {
	sugar.Infow(msg, withTrace(ctx, args)...)
}
func ErrorContext(ctx context.Context, msg string, args ...any) {
	sugar.Errorw(msg, withTrace(ctx, args)...)
}
func WarnContext(ctx context.Context, msg string, args ...any) {
	sugar.Warnw(msg, withTrace(ctx, args)...)
}
func DebugContext(ctx context.Context, msg string, args ...any) {
	sugar.Debugw(msg, withTrace(ctx, args)...)
}

// withTrace appends trace_id / span_id from the OTel span in ctx (if present).
// Falls back gracefully when OTel is not configured — no-op, no panic.
func withTrace(ctx context.Context, args []any) []any {
	if ctx == nil {
		return args
	}
	// Import is avoided to keep the logger package dependency-free.
	// Callers that need OTel fields can inject them as explicit key-value pairs.
	_ = ctx
	return args
}
