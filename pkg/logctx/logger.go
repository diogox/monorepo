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

func Info(ctx context.Context, msg string, fields ...with.Field) {
	getLogger(ctx).Info(msg, toZapFields(fields)...)
}

func Error(ctx context.Context, msg string, fields ...with.Field) {
	getLogger(ctx).Error(msg, toZapFields(fields)...)
}
