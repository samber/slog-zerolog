package slogzerolog

import (
	"context"

	"log/slog"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Option struct {
	// log level (default: debug)
	Level slog.Leveler

	// optional: zerolog logger (default: zerolog.Logger)
	Logger *zerolog.Logger

	// optional: customize json payload builder
	Converter Converter
}

func (o Option) NewZerologHandler() slog.Handler {
	if o.Level == nil {
		o.Level = slog.LevelDebug
	}

	if o.Logger == nil {
		// should be selected lazily ?
		o.Logger = &log.Logger
	}

	return &ZerologHandler{
		option: o,
		attrs:  []slog.Attr{},
		groups: []string{},
	}
}

var _ slog.Handler = (*ZerologHandler)(nil)

type ZerologHandler struct {
	option Option
	attrs  []slog.Attr
	groups []string
}

func (h *ZerologHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.option.Level.Level()
}

func (h *ZerologHandler) Handle(ctx context.Context, record slog.Record) error {
	converter := DefaultConverter
	if h.option.Converter != nil {
		converter = h.option.Converter
	}

	level := levelMap[record.Level]
	args := converter(h.attrs, &record)

	h.option.Logger.
		WithLevel(level).
		Ctx(ctx).
		Time(zerolog.TimestampFieldName, record.Time).
		Fields(args).
		Msg(record.Message)

	return nil
}

func (h *ZerologHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ZerologHandler{
		option: h.option,
		attrs:  appendAttrsToGroup(h.groups, h.attrs, attrs),
		groups: h.groups,
	}
}

func (h *ZerologHandler) WithGroup(name string) slog.Handler {
	return &ZerologHandler{
		option: h.option,
		attrs:  h.attrs,
		groups: append(h.groups, name),
	}
}