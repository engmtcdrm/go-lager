package lager

import (
	"context"
	"log/slog"
	"os"
	"strings"
)

func Init(logFileNm string, level slog.Leveler) (*os.File, error) {
	f, err := os.OpenFile(logFileNm, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}

	optsStderr := &HandlerOptions{
		Level:    level,
		AddLevel: true,
	}
	optsFile := &HandlerOptions{Level: level}

	handlerStdout := NewStdoutHandler(nil)
	handlerStderr := NewStderrHandler(optsStderr)
	handlerFile := NewFileHandler(f, optsFile)

	handlerMulti := NewMultiHandler(
		handlerStdout,
		handlerStderr,
		handlerFile,
	)

	logger := slog.New(handlerMulti)
	slog.SetDefault(logger)

	return f, nil
}

func Trace(msg string, args ...any) {
	trace(msg, args...)
}

func TraceIndent(msg string, indent int, args ...any) {
	trace(strings.Repeat(" ", indent)+msg, args...)
}

func DebugIndent(msg string, indent int, args ...any) {
	slog.Debug(strings.Repeat(" ", indent)+msg, args...)
}

func WarnIndent(msg string, indent int, args ...any) {
	slog.Warn(strings.Repeat(" ", indent)+msg, args...)
}

func ErrorIndent(msg string, indent int, args ...any) {
	slog.Error(strings.Repeat(" ", indent)+msg, args...)
}

func InfoIndent(msg string, indent int, args ...any) {
	slog.Info(strings.Repeat(" ", indent)+msg, args...)
}

func trace(msg string, args ...any) {
	slog.Default().Log(context.Background(), LevelTrace, msg, args...)
}
