package logging

import (
	"github.com/nanicienta/api/pkg/ports/logging"
	"go.uber.org/zap"
	"os"
)

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

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger(base *zap.Logger) *ZapLogger {
	return &ZapLogger{logger: base.Sugar()}
}

func (l *ZapLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

func (l *ZapLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

func (l *ZapLogger) Warn(msg string, keysAndValues ...interface{}) {
	l.logger.Warnw(msg, keysAndValues...)
}

func (l *ZapLogger) Error(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}

func (l *ZapLogger) Fatal(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues...)
}
