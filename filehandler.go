package lager

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
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
	w io.Writer
}

func NewFileHandler(filePath string) *FileHandler {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil
	}

	return &FileHandler{
		w: f,
	}
}

// Enabled checks if the handler is enabled for the given log level
func (h *FileHandler) Enabled(_ context.Context, level slog.Level) bool {
	return true
}

// Handle processes a log record and writes it to the file
func (h *FileHandler) Handle(_ context.Context, r slog.Record) error {
	msg := ansi.StripCodes(r.Message)
	timestamp := time.Now().Format("2006-01-02:15:04:05.000000")
	_, err := fmt.Fprintf(h.w, "%s [%-5s]: %s\n", timestamp, levelString(r.Level), msg)
	return err
}

// WithAttrs adds attributes to the handler
func (h *FileHandler) WithAttrs(attrs []slog.Attr) slog.Handler { return h }

// WithGroup adds a group name to the handler
func (h *FileHandler) WithGroup(name string) slog.Handler { return h }
