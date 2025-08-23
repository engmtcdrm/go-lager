package lager

import (
	"context"
	"fmt"
	"io"
	"log/slog"
)

// PlainHandler prints only the message text
type PlainHandler struct{ w io.Writer }

func NewPlainHandler(w io.Writer) *PlainHandler {
	return &PlainHandler{w: w}
}

// Enabled checks if the handler is enabled for the given log level
func (h *PlainHandler) Enabled(_ context.Context, _ slog.Level) bool { return true }

// Handle processes a log record and writes the message to the handler's output
func (h *PlainHandler) Handle(_ context.Context, r slog.Record) error {
	_, err := fmt.Fprintln(h.w, r.Message)
	return err
}

// WithAttrs returns a new handler with the given attributes added
func (h *PlainHandler) WithAttrs([]slog.Attr) slog.Handler { return h }

// WithGroup returns a new handler with the given group name
func (h *PlainHandler) WithGroup(string) slog.Handler { return h }
