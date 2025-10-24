package slogzerolog

import (
	"log/slog"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var LogLevels = map[slog.Level]zerolog.Level{
	slog.LevelDebug: zerolog.DebugLevel,
	slog.LevelInfo:  zerolog.InfoLevel,
	slog.LevelWarn:  zerolog.WarnLevel,
	slog.LevelError: zerolog.ErrorLevel,
}

var reverseLogLevels map[zerolog.Level]slog.Level

func init() {
	reverseLogLevels = map[zerolog.Level]slog.Level{}
	for level, z := range LogLevels {
		reverseLogLevels[z] = level
	}
}

// ZeroLogLeveler can be used for Option.Level (implements slog.Leveler).
type ZeroLogLeveler struct {
}

func (ZeroLogLeveler) Level() slog.Level {
	zeroLogLevel := log.Logger.GetLevel()
	level, ok := reverseLogLevels[zeroLogLevel]
	if !ok {
		return slog.LevelDebug
	}
	return level
}
