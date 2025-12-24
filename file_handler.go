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

func levelString(l slog.Level) string {
	if l < slog.LevelDebug {
		str := func(base string, val slog.Level) string {
			if val == 0 {
				return base
			}
			return fmt.Sprintf("%s%+d", base, val)
		}
		return str("TRACE", l-LevelTrace)
	}

	return l.String()
}

// FileHandler strips ANSI codes and prepends datetime
type FileHandler struct {
	w          io.Writer
	mu         *sync.Mutex
	opts       HandlerOptions
	timeFormat string
}

func NewFileHandler(w io.Writer, opts *HandlerOptions) *FileHandler {
	if opts == nil {
		opts = &HandlerOptions{}
	}

	if opts.Level == nil {
		opts.Level = LevelInfo
	}

	return &FileHandler{
		w:          w,
		mu:         &sync.Mutex{},
		opts:       *opts,
		timeFormat: time.RFC3339,
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
	msg := ansi.StripCodes(r.Message)
	var fullMsg string

	if r.Level == LevelInfo {
		fullMsg = fmt.Sprintf("%s - %s\n", r.Time.Round(0).Format(h.timeFormat), msg)
	} else {
		fullMsg = fmt.Sprintf("%s - %-5s: %s\n", r.Time.Round(0).Format(h.timeFormat), levelString(r.Level), msg)
	}

	// Create a buffer to hold the final output
	buf := make([]byte, 0, 1024)
	buf = append(buf, fullMsg...)

	// Write to the file
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.w.Write(buf)
	return err
}

// WithAttrs adds attributes to the handler
func (h *FileHandler) WithAttrs(attrs []slog.Attr) slog.Handler { return h }

// WithGroup adds a group name to the handler
func (h *FileHandler) WithGroup(name string) slog.Handler { return h }

func (h *FileHandler) TimeFormat(format string) *FileHandler {
	if format != "" {
		h.timeFormat = format
	}
	return h
}
