package lager

import (
	"context"
	"io"
	"log/slog"
	"sync"
)

// StreamHandler prints only the message text
type StreamHandler struct {
	w    io.Writer
	mu   *sync.Mutex
	opts Options
}

func NewStreamHandler(w io.Writer, opts *Options) *StreamHandler {
	h := &StreamHandler{
		w:  w,
		mu: &sync.Mutex{},
	}

	if opts != nil {
		h.opts = *opts
	}

	if h.opts.Level == nil {
		h.opts.Level = LevelInfo
	}
	return h
}

// Enabled checks if the handler is enabled for the given log level
func (h *StreamHandler) Enabled(ctx context.Context, level slog.Level) bool {
	if h.opts.Enablers == nil {
		return level >= h.opts.Level.Level()
	}

	for _, enable := range h.opts.Enablers {
		if !enable(ctx, level) {
			return false
		}
	}

	return true
}

// Handle processes a log record and writes the message to the handler's output
func (h *StreamHandler) Handle(ctx context.Context, r slog.Record) error {
	// Create a buffer to hold the final output
	buf := make([]byte, 0, 512)
	buf = append(buf, r.Message...)

	// Write to stream
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.w.Write(buf)
	return err
}

// WithAttrs returns a new handler with the given attributes added
func (h *StreamHandler) WithAttrs([]slog.Attr) slog.Handler { return h }

// WithGroup returns a new handler with the given group name
func (h *StreamHandler) WithGroup(string) slog.Handler { return h }
