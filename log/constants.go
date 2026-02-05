package log

import (
	"log/slog"
	"slices"
)

const (

	// custom log levels
	levelTrace slog.Level = -8

	// colors
	colorReset   = "\033[0m"
	colorMagenta = "\033[35m"
	colorBlue    = "\033[34m"
	colorYellow  = "\033[33m"
	colorRed     = "\033[31m"
	colorCyan    = "\033[36m"
	colorWhite   = "\033[37m"
)

var (
	LOG_TOP_LEVEL_HIDDEN_FIELDS = []string{
		"stage",
		"phase",
	}
)

func init() {
	slices.Sort(LOG_TOP_LEVEL_HIDDEN_FIELDS)
}
