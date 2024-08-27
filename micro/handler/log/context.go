package log

import (
	"context"

	"go-micro.dev/v4/logger"
)

type loggerContextKey string

var defaultLoggerContextKey = loggerContextKey("ctx-logger")

func LoggerFromContext(ctx context.Context) *logger.Helper {
	if hl, ok := ctx.Value(defaultLoggerContextKey).(*logger.Helper); ok {
		return hl
	}
	return nil
}
