package logctx

import (
	"context"

	"github.com/diogox/monorepo/pkg/logctx/with"
)

type ctxKey string

const ctxKeyLogger = ctxKey("logger")

// Propagate returns a new context.Context that, when used in future logs, will include the provided log fields.
func Propagate(ctx context.Context, field with.Field, otherFields ...with.Field) context.Context {
	zapLogger := getLogger(ctx).
		Logger.
		With(field()).
		With(toZapFields(otherFields)...)

	return toCtx(ctx, ctxLogger{
		Logger: zapLogger,
	})
}

func getLogger(ctx context.Context) ctxLogger {
	logger, ok := fromCtx(ctx)
	if !ok {
		return defaultCtxLogger
	}

	return logger
}

func fromCtx(ctx context.Context) (ctxLogger, bool) {
	logger, ok := ctx.Value(ctxKeyLogger).(ctxLogger)
	if !ok {
		return ctxLogger{}, false
	}

	return logger, true
}

func toCtx(ctx context.Context, logger ctxLogger) context.Context {
	return context.WithValue(ctx, ctxKeyLogger, logger)
}
