package logger

import (
	"context"
	"log/slog"
	"os"
)

const (
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	reset  = "\033[0m"
)

// ColorsHandler envuelve el TextHandler y agrega colores según el nivel del log
type ColorsHandler struct {
	slog.Handler
}

func NewColorsHandler(h slog.Handler) *ColorsHandler {
	return &ColorsHandler{Handler: h}
}

func (h *ColorsHandler) Handle(ctx context.Context, r slog.Record) error {
	var color string
	switch r.Level {
	case slog.LevelError:
		color = red
	case slog.LevelWarn:
		color = yellow
	case slog.LevelInfo:
		color = green
	default:
		color = reset
	}

	os.Stdout.Write([]byte(color))
	err := h.Handler.Handle(ctx, r)
	os.Stdout.Write([]byte(reset))
	return err
}

func (h *ColorsHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewColorsHandler(h.Handler.WithAttrs(attrs))
}
