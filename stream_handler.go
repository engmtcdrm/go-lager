package lager

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/engmtcdrm/go-ansi"
)

type StreamType int

const (
	StreamStdout StreamType = iota
	StreamStderr
)

// streamHandler prints only the message text
type streamHandler struct {
	w    io.Writer
	mu   *sync.Mutex
	opts HandlerOptions
}

// newStreamHandler creates a new StreamHandler with the given writer and options
func newStreamHandler(st StreamType, opts *HandlerOptions) *streamHandler {
	if opts == nil {
		opts = &HandlerOptions{}
	}

	var w io.Writer
	if st == StreamStdout {
		w = os.Stdout
	} else {
		w = os.Stderr
	}

	return &streamHandler{
		w:    w,
		mu:   &sync.Mutex{},
		opts: *opts,
	}
}

// Enabled checks if the handler is enabled for the given log level
func (h *streamHandler) Enabled(ctx context.Context, level slog.Level) bool {
	minLevel := LevelInfo
	if h.opts.Level != nil {
		minLevel = h.opts.Level.Level()
	}

	if h.opts.Enablers == nil {
		return level >= minLevel
	}

	for _, enable := range h.opts.Enablers {
		if !enable(ctx, level) {
			return false
		}
	}

	return true
}

// Handle processes a log record and writes the message to the handler's output
func (h *streamHandler) Handle(ctx context.Context, r slog.Record) error {
	// Create a buffer to hold the final output
	buf := make([]byte, 0, 512)

	if h.opts.AddTime && !r.Time.IsZero() {
		buf = append(buf, r.Time.Round(0).Format(time.RFC3339)...)
		buf = append(buf, " - "...)
	}

	if h.opts.AddLevel {
		buf = append(buf, levelString(r.Level)...)
		buf = append(buf, ": "...)
	}

	if h.opts.AddSource && r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		buf = append(buf, slog.String(slog.SourceKey, fmt.Sprintf("%s:%d", f.File, f.Line)).Value.String()...)
		buf = append(buf, " - "...)
	}

	buf = append(buf, r.Message+"\n"...)

	if h.opts.NoColor {
		buf = []byte(ansi.StripCodes(string(buf)))
	}

	// Write to stream
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.w.Write(buf)
	return err
}

// WithAttrs returns a new handler with the given attributes added
func (h *streamHandler) WithAttrs([]slog.Attr) slog.Handler {
	return h
}

// WithGroup returns a new handler with the given group name
func (h *streamHandler) WithGroup(string) slog.Handler {
	return h
}
