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

// logDebug controls whether debug messages are logged
// var logDebug = false

func Init(logFileNm string, debug bool) (*os.File, error) {
	// logDebug = debug

	f, err := os.OpenFile(logFileNm, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}

	handlerStdout := NewStdoutHandler(nil)
	handlerStderr := NewStderrHandler(nil)
	handlerFile := NewFileHandler(f, nil)

	handlerMulti := NewMultiHandler(
		handlerStdout,
		handlerStderr,
		handlerFile,
	)

	logger := slog.New(handlerMulti)
	slog.SetDefault(logger)

	return f, nil
}

func Debug(msg string, args ...any) {
	slog.Debug(msg, args...)
}

func DebugIndent(msg string, indent int, args ...any) {
	slog.Debug(strings.Repeat(" ", indent)+msg, args...)
}

func Error(msg string, args ...any) {
	slog.Error(msg, args...)
}

func ErrorIndent(msg string, indent int, args ...any) {
	slog.Error(strings.Repeat(" ", indent)+msg, args...)
}

func Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

func InfoIndent(msg string, indent int, args ...any) {
	slog.Info(strings.Repeat(" ", indent)+msg, args...)
}

func trace(msg string, args ...any) {
	slog.Default().Log(context.Background(), LevelTrace, msg, args...)
}

func Trace(msg string, args ...any) {
	trace(msg, args...)
}

func TraceIndent(msg string, indent int, args ...any) {
	trace(strings.Repeat(" ", indent)+msg, args...)
}

func Warn(msg string, args ...any) {
	slog.Warn(msg, args...)
}

func WarnIndent(msg string, indent int, args ...any) {
	slog.Warn(strings.Repeat(" ", indent)+msg, args...)
}
