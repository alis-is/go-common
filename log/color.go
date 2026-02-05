package log

import "log/slog"

func levelName(level slog.Level) string {
	switch {
	case level < slog.LevelDebug:
		return "TRACE"
	case level < slog.LevelInfo:
		return "DEBUG"
	case level < slog.LevelWarn:
		return "INFO"
	case level < slog.LevelError:
		return "WARNING"
	default:
		return "ERROR"
	}
}

func levelColor(level slog.Level) string {
	switch {
	case level < slog.LevelDebug:
		return colorCyan
	case level < slog.LevelInfo:
		return colorMagenta
	case level < slog.LevelWarn:
		return colorBlue
	case level < slog.LevelError:
		return colorYellow
	default:
		return colorRed
	}
}

func colorize(value string, color string, enabled bool) string {
	if !enabled || color == "" {
		return value
	}
	return color + value + colorReset
}
