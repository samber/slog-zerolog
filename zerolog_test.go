package slogzerolog

import (
	"log/slog"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestZeroLogLeveler_Level(t *testing.T) {
	type fields struct {
		Logger *zerolog.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   slog.Level
	}{
		{
			name: "DebugLevel",
			fields: fields{
				Logger: toPtr(log.Logger.Level(zerolog.DebugLevel)),
			},
			want: slog.LevelDebug,
		},
		{
			name: "InfoLevel",
			fields: fields{
				Logger: toPtr(log.Logger.Level(zerolog.InfoLevel)),
			},
			want: slog.LevelInfo,
		},
		{
			name: "WarnLevel",
			fields: fields{
				Logger: toPtr(log.Logger.Level(zerolog.WarnLevel)),
			},
			want: slog.LevelWarn,
		},
		{
			name: "ErrorLevel",
			fields: fields{
				Logger: toPtr(log.Logger.Level(zerolog.ErrorLevel)),
			},
			want: slog.LevelError,
		},
		{
			name: "FatalLevel",
			fields: fields{
				Logger: toPtr(log.Logger.Level(zerolog.FatalLevel)),
			},
			want: slog.LevelError,
		},
		{
			name: "PanicLevel",
			fields: fields{
				Logger: toPtr(log.Logger.Level(zerolog.PanicLevel)),
			},
			want: slog.LevelError,
		},
		{
			name: "NoLevel",
			fields: fields{
				Logger: toPtr(log.Logger.Level(zerolog.NoLevel)),
			},
			want: slog.LevelInfo,
		},
		{
			name: "TraceLevel",
			fields: fields{
				Logger: toPtr(log.Logger.Level(zerolog.TraceLevel)),
			},
			want: slog.Level(-5),
		},
		{
			name: "CustomTraceLevel",
			fields: fields{
				Logger: toPtr(log.Logger.Level(-42)),
			},
			want: slog.Level(-46),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := ZeroLogLeveler{
				Logger: tt.fields.Logger,
			}
			if got := z.Level(); got != tt.want {
				t.Errorf("Level() = %v, want %v", got, tt.want)
			}
		})
	}
}

func toPtr[T any](v T) *T { return &v }
