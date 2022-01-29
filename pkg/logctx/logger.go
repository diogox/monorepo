package logctx

import (
	"context"
	"github.com/diogox/monorepo/pkg/logctx/with"

	"go.uber.org/zap"
)

type ctxLogger struct {
	*zap.Logger
}

var defaultCtxLogger ctxLogger

func init() {
	// TODO: Look into providing more options. Like syntactic sugar and logging levels.

	logger, err := zap.NewProduction()
	if err != nil {
		panic("failed to initialize default context logger")
	}

	defaultCtxLogger = ctxLogger{
		logger,
	}
}

func toZapFields(ff []with.Field) []zap.Field {
	zapFields := make([]zap.Field, len(ff))
	for i, f := range ff {
		zapFields[i] = f()
	}

	return zapFields
}

// Info logs an informative message.
func Info(ctx context.Context, msg string, fields ...with.Field) {
	getLogger(ctx).Info(msg, toZapFields(fields)...)
}

// Error logs an error message.
func Error(ctx context.Context, msg string, fields ...with.Field) {
	getLogger(ctx).Error(msg, toZapFields(fields)...)
}

// Debug logs a message for debugging purposes.
func Debug(ctx context.Context, msg string, fields ...with.Field) {
	getLogger(ctx).Debug(msg, toZapFields(fields)...)
}
