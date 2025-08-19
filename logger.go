package lager

import (
	"context"
	"log/slog"
	"os"
	"strings"

	slogmulti "github.com/samber/slog-multi"
)

const LevelTrace slog.Level = -8

// logDebug controls whether debug messages are logged
var logDebug = false

func Init(logFileNm string, debug bool) *os.File {
	logDebug = debug

	// Open log file
	logFile, err := os.OpenFile(logFileNm, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	// Define various logging routs
	router := slogmulti.Router().
		Add(&PlainHandler{os.Stdout}, func(_ context.Context, r slog.Record) bool {
			return r.Level == slog.LevelInfo
		}).
		Add(&PlainHandler{os.Stderr}, func(_ context.Context, r slog.Record) bool {
			if !logDebug && r.Level == slog.LevelDebug {
				return false
			}
			return r.Level != slog.LevelInfo
		}).
		Add(&FileHandler{w: logFile}, func(_ context.Context, r slog.Record) bool {
			if !logDebug && r.Level == slog.LevelDebug {
				return false
			}
			return true
		})

	logger := slog.New(router.Handler())
	slog.SetDefault(logger)

	return logFile
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
