package slogzerolog

import (
	"context"

	"log/slog"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	slogcommon "github.com/samber/slog-common"
)

type Option struct {
	// log level (default: debug)
	Level slog.Leveler

	// optional: zerolog logger (default: zerolog.Logger)
	Logger *zerolog.Logger

	// optional: customize json payload builder
	Converter Converter
	// optional: fetch attributes from context
	AttrFromContext []func(ctx context.Context) []slog.Attr

	// optional: see slog.HandlerOptions
	AddSource   bool
	ReplaceAttr func(groups []string, a slog.Attr) slog.Attr
}

func (o Option) NewZerologHandler() slog.Handler {
	if o.Level == nil {
		o.Level = slog.LevelDebug
	}

	if o.Logger == nil {
		// should be selected lazily ?
		o.Logger = &log.Logger
	}

	if o.AttrFromContext == nil {
		o.AttrFromContext = []func(ctx context.Context) []slog.Attr{}
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

	level := LogLevels[record.Level]
	fromContext := slogcommon.ContextExtractor(ctx, h.option.AttrFromContext)
	args := converter(h.option.AddSource, h.option.ReplaceAttr, append(h.attrs, fromContext...), h.groups, &record)

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
		attrs:  slogcommon.AppendAttrsToGroup(h.groups, h.attrs, attrs...),
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
