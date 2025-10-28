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
		opts = &Options{}
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
func (h *StdoutHandler) WithAttrs([]slog.Attr) slog.Handler { return h }

// WithGroup returns a new handler with the given group name
func (h *StdoutHandler) WithGroup(string) slog.Handler { return h }
