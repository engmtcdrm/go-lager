package lager

import (
	"context"
	"log/slog"
)

type StdoutHandler struct {
	streamHandler streamHandler
}

func NewStdoutHandler(opts *HandlerOptions) *StdoutHandler {
	if opts == nil {
		opts = &HandlerOptions{}
	}

	if opts.Level == nil {
		opts.Level = slog.LevelInfo
	}

	if opts.Enablers == nil {
		opts.Enablers = []func(ctx context.Context, level slog.Level) bool{}
	}

	opts.Enablers = append(opts.Enablers, func(ctx context.Context, level slog.Level) bool {
		return level == slog.LevelInfo
	})

	sh := newStreamHandler(StreamStdout, opts)

	return &StdoutHandler{
		streamHandler: *sh,
	}
}

// Enabled checks if the handler is enabled for the given log level
func (h *StdoutHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.streamHandler.Enabled(ctx, level)
}

// Handle processes a log record and writes the message to the handler's output
func (h *StdoutHandler) Handle(ctx context.Context, r slog.Record) error {
	return h.streamHandler.Handle(ctx, r)
}

// WithAttrs returns a new handler with the given attributes added
func (h *StdoutHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h.streamHandler.WithAttrs(attrs)
}

// WithGroup returns a new handler with the given group name
func (h *StdoutHandler) WithGroup(name string) slog.Handler {
	return h.streamHandler.WithGroup(name)
}
