package logger

import (
	"github.com/ascend/phoenix-multichannel-mkp/config"
	"github.com/hermese-team/phoenix-multichannel-mkp/libraries/go/telemetry"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.uber.org/zap"
)

// global is the package-level sugared logger. Defaults to a no-op so service
// code can safely call Error/Info before New is called.
var global = zap.NewNop().Sugar()

// New creates a *zap.Logger, sets the package-level global, and returns both
// so the caller can defer log.Sync() and lp.Shutdown(ctx).
func New(cfg *config.Config) (*zap.Logger, *sdklog.LoggerProvider, error) {
	log, lp, err := telemetry.NewLogger(telemetry.Options{
		Endpoint:    cfg.Telemetry.Endpoint,
		ServiceName: cfg.Telemetry.ServiceName,
		Env:         cfg.App.Env,
		LogLevel:    cfg.App.LogLevel,
	})
	if err != nil {
		return nil, nil, err
	}
	global = log.Sugar()
	return log, lp, nil
}

// Debug logs at DEBUG level with slog-style key-value pairs.
func Debug(msg string, keysAndValues ...any) { global.Debugw(msg, keysAndValues...) }

// Info logs at INFO level with slog-style key-value pairs.
func Info(msg string, keysAndValues ...any) { global.Infow(msg, keysAndValues...) }

// Warn logs at WARN level with slog-style key-value pairs.
func Warn(msg string, keysAndValues ...any) { global.Warnw(msg, keysAndValues...) }

// Error logs at ERROR level with slog-style key-value pairs.
func Error(msg string, keysAndValues ...any) { global.Errorw(msg, keysAndValues...) }
