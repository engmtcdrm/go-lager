package lager

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"sync"
	"time"

	"github.com/engmtcdrm/go-ansi"
)

// FileHandler is a handler that writes log messages to a specified file or
// io.Writer, stripping ANSI escape codes from the messages.
type FileHandler struct {
	w    io.Writer
	mu   *sync.Mutex
	opts HandlerOptions
}

// NewFileHandler creates a new [FileHandler] with the given [io.Writer] and [HandlerOptions].
func NewFileHandler(w io.Writer, opts *HandlerOptions) *FileHandler {
	if opts == nil {
		opts = &HandlerOptions{}
	}

	if opts.Level == nil {
		opts.Level = LevelInfo
	}

	if opts.TimeFormat == "" {
		opts.TimeFormat = time.RFC3339
	}

	return &FileHandler{
		w:    w,
		mu:   &sync.Mutex{},
		opts: *opts,
	}
}

// Enabled checks if the handler is enabled for the given log level
func (h *FileHandler) Enabled(ctx context.Context, level slog.Level) bool {
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

// Handle processes a log record and writes it to the file
func (h *FileHandler) Handle(ctx context.Context, r slog.Record) error {
	// Create a buffer to hold the final output
	buf := make([]byte, 0, 1024)

	if h.opts.AddTime && !r.Time.IsZero() {
		buf = append(buf, r.Time.Round(0).Format(h.opts.TimeFormat)...)
		buf = append(buf, " - "...)
	}

	if h.opts.AddLevel {
		buf = append(buf, fmt.Sprintf("%-5s - ", levelString(r.Level))...)
	}

	if h.opts.AddSource && r.Level < LevelInfo {
		src := r.Source()
		if src == nil {
			src = &slog.Source{}
		}
		buf = append(buf, slog.String(slog.SourceKey, fmt.Sprintf("%s:%d", src.File, src.Line)).Value.String()...)
		buf = append(buf, " - "...)
	}

	msg := ansi.StripCodes(r.Message)
	buf = append(buf, msg+"\n"...)

	// Write to the file
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.w.Write(buf)
	return err
}

// WithAttrs adds attributes to the handler
func (h *FileHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

// WithGroup adds a group name to the handler
func (h *FileHandler) WithGroup(name string) slog.Handler {
	return h
}
