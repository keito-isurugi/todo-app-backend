package logger

import (
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type sentryCore struct {
	zapcore.Core
}

func (c *sentryCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	if entry.Level == zapcore.ErrorLevel {
		sentry.CaptureException(errors.New(entry.Message))
	}
	return c.Core.Write(entry, fields)
}

func (c *sentryCore) With(fields []zapcore.Field) zapcore.Core {
	return &sentryCore{c.Core.With(fields)}
}

func NewSentryCore(baseCore zapcore.Core) zapcore.Core {
	return &sentryCore{baseCore}
}

func NewLogger(debug bool) (*zap.Logger, error) {
	logLevel := zap.InfoLevel
	if debug {
		logLevel = zap.DebugLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		logLevel,
	)

	sentryCore := NewSentryCore(core)
	return zap.New(sentryCore), nil
}
