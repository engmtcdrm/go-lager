package lager

import (
	"context"
	"log/slog"
	"os"
	"strings"
)

const (
	LevelTrace   = slog.Level(-8)
	LevelDebug   = slog.LevelDebug
	LevelInfo    = slog.LevelInfo
	LevelWarning = slog.LevelWarn
	LevelError   = slog.LevelError
)

func trace(msg string, args ...any) {
	slog.Default().Log(context.Background(), LevelTrace, msg, args...)
}

func Init(logFileNm string, level slog.Leveler) (*os.File, error) {
	f, err := os.OpenFile(logFileNm, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}

	opts := &HandlerOptions{Level: level}
	opts2 := *opts

	handlerStdout := NewStdoutHandler(nil)
	handlerStderr := NewStderrHandler(opts)
	handlerFile := NewFileHandler(f, &opts2)

	handlerMulti := NewMultiHandler(
		handlerStdout,
		handlerStderr,
		handlerFile,
	)

	logger := slog.New(handlerMulti)
	slog.SetDefault(logger)

	return f, nil
}

func DebugIndent(msg string, indent int, args ...any) {
	slog.Debug(strings.Repeat(" ", indent)+msg, args...)
}

func ErrorIndent(msg string, indent int, args ...any) {
	slog.Error(strings.Repeat(" ", indent)+msg, args...)
}

func InfoIndent(msg string, indent int, args ...any) {
	slog.Info(strings.Repeat(" ", indent)+msg, args...)
}

func Trace(msg string, args ...any) {
	trace(msg, args...)
}

func TraceIndent(msg string, indent int, args ...any) {
	trace(strings.Repeat(" ", indent)+msg, args...)
}

func WarnIndent(msg string, indent int, args ...any) {
	slog.Warn(strings.Repeat(" ", indent)+msg, args...)
}
