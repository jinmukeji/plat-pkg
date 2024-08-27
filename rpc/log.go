package rpc

import (
	"context"

	"github.com/sirupsen/logrus"
)

type loggerContextKey struct{}

func ContextWithLogger(ctx context.Context, l *logrus.Entry) context.Context {
	c := context.WithValue(ctx, loggerContextKey{}, l)
	return c
}

func LoggerFromContext(ctx context.Context) *logrus.Entry {
	if l, ok := ctx.Value(loggerContextKey{}).(*logrus.Entry); ok {
		return l
	}

	return nil
}
