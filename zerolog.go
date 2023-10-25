package slogzerolog

import (
	"log/slog"

	"github.com/rs/zerolog"
)

var LogLevels = map[slog.Level]zerolog.Level{
	slog.LevelDebug: zerolog.DebugLevel,
	slog.LevelInfo:  zerolog.InfoLevel,
	slog.LevelWarn:  zerolog.WarnLevel,
	slog.LevelError: zerolog.ErrorLevel,
}
