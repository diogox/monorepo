package logctx

import (
	"context"
)

type ctxKey string

const ctxKeyLogger = ctxKey("logger")

func getLogger(ctx context.Context) ctxLogger {
	logger, ok := fromCtx(ctx)
	if !ok {
		logger = defaultCtxLogger
	}

	toCtx(ctx, logger)

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
