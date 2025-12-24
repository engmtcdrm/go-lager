package lager

import (
	"context"
	"log/slog"
)

type StderrHandler struct {
	streamHandler StreamHandler
}

func NewStderrHandler(opts *HandlerOptions) *StderrHandler {
	if opts == nil {
		opts = &HandlerOptions{}
	}

	if opts.Level == nil {
		opts.Level = slog.LevelWarn
	}

	if opts.Enablers == nil {
		opts.Enablers = []func(ctx context.Context, level slog.Level) bool{}
	}

	// Never write info level to stderr
	opts.Enablers = append(opts.Enablers, func(ctx context.Context, level slog.Level) bool {
		return level != LevelInfo
	})

	// Based on the specified log level, add appropriate enablers
	switch opts.Level {
	case LevelDebug:
		opts.Enablers = append(opts.Enablers, func(ctx context.Context, level slog.Level) bool {
			return level >= LevelDebug
		})
	case LevelTrace:
		opts.Enablers = append(opts.Enablers, func(ctx context.Context, level slog.Level) bool {
			return level >= LevelTrace
		})
	default:
		opts.Enablers = append(opts.Enablers, func(ctx context.Context, level slog.Level) bool {
			return level >= LevelWarning
		})
	}

	sh := NewStreamHandler(StreamStderr, opts)

	return &StderrHandler{
		streamHandler: *sh,
	}
}

// Enabled checks if the handler is enabled for the given log level
func (h *StderrHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.streamHandler.Enabled(ctx, level)
}

// Handle processes a log record and writes the message to the handler's output
func (h *StderrHandler) Handle(ctx context.Context, r slog.Record) error {
	return h.streamHandler.Handle(ctx, r)
}

// WithAttrs returns a new handler with the given attributes added
func (h *StderrHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h.streamHandler.WithAttrs(attrs)
}

// WithGroup returns a new handler with the given group name
func (h *StderrHandler) WithGroup(name string) slog.Handler {
	return h.streamHandler.WithGroup(name)
}
