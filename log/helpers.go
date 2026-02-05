package log

import (
	"context"
	"log/slog"
	"slices"
	"strings"
)

func Trace(msg string, args ...any) {
	slog.Log(context.Background(), levelTrace, msg, args...)
}

func Debug(msg string, args ...any) {
	slog.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	slog.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	slog.Error(msg, args...)
}

func ParseLevel(levelFlag string) (slog.Level, string) {
	switch strings.ToLower(levelFlag) {
	case "trace":
		return levelTrace, "trace"
	case "debug":
		return slog.LevelDebug, "debug"
	case "warn", "warning":
		return slog.LevelWarn, "warn"
	case "error":
		return slog.LevelError, "error"
	default:
		return slog.LevelInfo, "info"
	}
}

func isHiddenAttr(attr slog.Attr) bool {
	_, found := slices.BinarySearch(LOG_TOP_LEVEL_HIDDEN_FIELDS, attr.Key)
	return found
}
