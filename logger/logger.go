package logger

import (
	"context"
	"log/slog"

	"github.com/fulldump/box"
	"github.com/google/uuid"
)

func GetLog(ctx context.Context) *slog.Logger {
	return ctx.Value("log").(*slog.Logger)
}

func InjectLog(log *slog.Logger) box.I {
	return func(next box.H) box.H {
		return func(ctx context.Context) {

			r := box.GetRequest(ctx)

			traceparent := r.Header.Get("Traceparent")
			if traceparent == "" {
				traceparent = uuid.NewString()
			}

			l := log.With(
				"traceparent", traceparent,
				"action", box.GetBoxContext(ctx).Action.Name,
			)
			ctx = SetLog(ctx, l)

			w := box.GetResponse(ctx)
			w.Header().Set("Traceparent", traceparent)

			next(ctx)
		}
	}
}

func SetLog(ctx context.Context, log *slog.Logger) context.Context {
	ctx = context.WithValue(ctx, "log", log)
	return ctx
}
