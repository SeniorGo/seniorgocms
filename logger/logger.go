package logger

import (
	"context"
	"github.com/fulldump/box"
	"log/slog"
)

func GetLog(ctx context.Context) *slog.Logger {
	return ctx.Value("log").(*slog.Logger)
}

func InjectLog(log *slog.Logger) box.I {
	return func(next box.H) box.H {
		return func(ctx context.Context) {
			ctx = SetLog(ctx, log)
			next(ctx)
		}
	}
}

func SetLog(ctx context.Context, log *slog.Logger) context.Context {
	ctx = context.WithValue(ctx, "log", log)
	return ctx
}
