// Package logging provides interfaces and implementations for structured and contextual logging.
package logging

import (
	"github.com/nanicienta/api/pkg/ports/logging"
	"go.uber.org/zap"
	"os"
)

// InitZapLogger initializes a zap logger with the production configuration.
func InitZapLogger() logging.Logger {
	cfg := zap.NewProductionConfig()
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	_ = cfg.Level.UnmarshalText([]byte(logLevel))

	zapLogger, _ := cfg.Build()
	return NewZapLogger(zapLogger)
}

// ZapLogger is a wrapper around zap.SugaredLogger to implement the logging.Logger interface.
type ZapLogger struct {
	logger *zap.SugaredLogger
}

// NewZapLogger creates a new ZapLogger instance.
func NewZapLogger(base *zap.Logger) *ZapLogger {
	return &ZapLogger{logger: base.Sugar()}
}

// Debug logs a message at debug level.
func (l *ZapLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

// Info logs a message at info level.
func (l *ZapLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

// Warn logs a message at warn level.
func (l *ZapLogger) Warn(msg string, keysAndValues ...interface{}) {
	l.logger.Warnw(msg, keysAndValues...)
}

// Error logs a message at error level.
func (l *ZapLogger) Error(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}

// Fatal logs a message at panic level.
func (l *ZapLogger) Fatal(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues...)
}
