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
		enablers := []func(ctx context.Context, level slog.Level) bool{}

		enablers = append(enablers, func(ctx context.Context, level slog.Level) bool {
			return level == slog.LevelInfo
		})

		opts = &Options{
			Enablers: enablers,
		}
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
	h.streamHandler.mu.Lock()
	defer h.streamHandler.mu.Unlock()
	_, err := fmt.Fprintf(h.streamHandler.w, "%s\n", r.Message)
	return err
}

// WithAttrs returns a new handler with the given attributes added
func (h *StderrHandler) WithAttrs([]slog.Attr) slog.Handler { return h }

// WithGroup returns a new handler with the given group name
func (h *StderrHandler) WithGroup(string) slog.Handler { return h }
