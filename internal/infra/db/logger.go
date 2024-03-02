package db

import (
	"context"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

const (
	SlowThresholdDuration = 100 * time.Millisecond
)

type GormLogger struct {
	*zap.Logger
	LogLevel         logger.LogLevel
	SlowThreshold    time.Duration
	SkipCallerLookup bool
}

func initGormLogger(zapLogger *zap.Logger) *GormLogger {
	return &GormLogger{
		Logger:           zapLogger,
		LogLevel:         logger.Info,
		SlowThreshold:    SlowThresholdDuration,
		SkipCallerLookup: false,
	}
}

func (l *GormLogger) logger() *zap.Logger {
	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		default:
			return l.Logger.WithOptions(zap.AddCallerSkip(i))
		}
	}
	return l.Logger
}

// Warn log Warn for Gorm
func (l *GormLogger) Warn(_ context.Context, str string, args ...interface{}) {
	if l.LogLevel < logger.Warn {
		return
	}
	l.logger().Sugar().Warnf(str, args...)
}
