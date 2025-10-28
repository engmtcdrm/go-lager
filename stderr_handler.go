package lager

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

type StderrHandler struct {
	streamHandler StreamHandler
}

func NewStderrHandler(opts *Options) *StderrHandler {
	if opts == nil {
		opts = &Options{}
	}

	if opts.Level == nil {
		opts.Level = slog.LevelWarn
	}

	if opts.Enablers == nil {
		opts.Enablers = []func(ctx context.Context, level slog.Level) bool{}
	}

	opts.Enablers = append(opts.Enablers, func(ctx context.Context, level slog.Level) bool {
		return level != LevelInfo
	})

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

	h := NewStreamHandler(os.Stderr, opts)

	h2 := &StderrHandler{
		streamHandler: *h,
	}

	return h2
}

// Enabled checks if the handler is enabled for the given log level
func (h *StderrHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.streamHandler.Enabled(ctx, level)
}

// Handle processes a log record and writes the message to the handler's output
func (h *StderrHandler) Handle(ctx context.Context, r slog.Record) error {
	fullMsg := fmt.Sprintf("%s\n", r.Message)
	// Create a buffer to hold the final output
	buf := make([]byte, 0, 512)
	buf = append(buf, fullMsg...)

	// Write to the file
	h.streamHandler.mu.Lock()
	defer h.streamHandler.mu.Unlock()
	_, err := h.streamHandler.w.Write(buf)
	return err
}

// WithAttrs returns a new handler with the given attributes added
func (h *StderrHandler) WithAttrs([]slog.Attr) slog.Handler { return h }

// WithGroup returns a new handler with the given group name
func (h *StderrHandler) WithGroup(string) slog.Handler { return h }
