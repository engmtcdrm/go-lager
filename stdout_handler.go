package lager

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

type StdoutHandler struct {
	streamHandler StreamHandler
}

func NewStdoutHandler(opts *Options) *StdoutHandler {
	if opts == nil {
		enablers := []func(ctx context.Context, level slog.Level) bool{}

		enablers = append(enablers, func(ctx context.Context, level slog.Level) bool {
			return level != slog.LevelInfo
		})

		opts = &Options{
			Enablers: enablers,
		}
	}

	h := NewStreamHandler(os.Stdout, opts)

	h2 := &StdoutHandler{
		streamHandler: *h,
	}

	return h2
}

// Enabled checks if the handler is enabled for the given log level
func (h *StdoutHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.streamHandler.Enabled(ctx, level)
}

// Handle processes a log record and writes the message to the handler's output
func (h *StdoutHandler) Handle(ctx context.Context, r slog.Record) error {
	h.streamHandler.mu.Lock()
	defer h.streamHandler.mu.Unlock()
	_, err := fmt.Fprintf(h.streamHandler.w, "%s\n", r.Message)
	return err
}

// WithAttrs returns a new handler with the given attributes added
func (h *StdoutHandler) WithAttrs([]slog.Attr) slog.Handler { return h }

// WithGroup returns a new handler with the given group name
func (h *StdoutHandler) WithGroup(string) slog.Handler { return h }
