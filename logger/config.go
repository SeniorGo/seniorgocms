package logger

import (
	"log/slog"
	"strings"
)

type ConfigLog struct {
	Type   string `json:"type"`
	Colors bool   `json:"colors"`
	Level  string `json:"level"`
}

var levels = map[string]slog.Level{
	"debug":   slog.LevelDebug,
	"info":    slog.LevelInfo,
	"warn":    slog.LevelWarn,
	"warning": slog.LevelWarn,
	"error":   slog.LevelError,
}

func ParseLevel(level string) slog.Level {
	level = strings.ToLower(level) // normalize
	v, ok := levels[level]
	if !ok {
		return slog.LevelInfo
	}
	return v
}
