package logger

import (
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

func Info(msg string, args ...any)  { sugar.Infow(msg, args...) }
func Error(msg string, args ...any) { sugar.Errorw(msg, args...) }
func Warn(msg string, args ...any)  { sugar.Warnw(msg, args...) }
func Debug(msg string, args ...any) { sugar.Debugw(msg, args...) }
