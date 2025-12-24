package lager

import (
	"context"
	"log/slog"
	"strings"
)

// IndentSpaces defines the number of spaces per indentation level. Default is 2.
var IndentSpaces = uint(2)

// Trace calls [LevelTrace] on the default logger.
func Trace(msg string, args ...any) {
	trace(msg, args...)
}

// TraceContext calls [LevelTrace] on the default logger with context.
func TraceContext(ctx context.Context, msg string, args ...any) {
	traceContext(ctx, msg, args...)
}

// TraceIndent calls [trace] on the default logger with indentation.
func TraceIndent(msg string, indent int, args ...any) {
	trace(strings.Repeat(" ", indent*int(IndentSpaces))+msg, args...)
}

// TraceContextIndent calls [traceContext] on the default logger with indentation.
func TraceContextIndent(ctx context.Context, msg string, indent int, args ...any) {
	traceContext(ctx, strings.Repeat(" ", indent*int(IndentSpaces))+msg, args...)
}

// DebugIndent calls [slog.Debug] on the default logger with indentation.
func DebugIndent(msg string, indent int, args ...any) {
	slog.Debug(strings.Repeat(" ", indent*int(IndentSpaces))+msg, args...)
}

// DebugContextIndent calls [slog.DebugContext] on the default logger with indentation.
func DebugContextIndent(ctx context.Context, msg string, indent int, args ...any) {
	slog.DebugContext(ctx, strings.Repeat(" ", indent*int(IndentSpaces))+msg, args...)
}

// WarnIndent calls [slog.Warn] on the default logger with indentation.
func WarnIndent(msg string, indent int, args ...any) {
	slog.Warn(strings.Repeat(" ", indent*int(IndentSpaces))+msg, args...)
}

// WarnContextIndent calls [slog.WarnContext] on the default logger with indentation.
func WarnContextIndent(ctx context.Context, msg string, indent int, args ...any) {
	slog.WarnContext(ctx, strings.Repeat(" ", indent*int(IndentSpaces))+msg, args...)
}

// ErrorIndent calls [slog.Error] on the default logger with indentation.
func ErrorIndent(msg string, indent int, args ...any) {
	slog.Error(strings.Repeat(" ", indent*int(IndentSpaces))+msg, args...)
}

// ErrorContextIndent calls [slog.ErrorContext] on the default logger with indentation.
func ErrorContextIndent(ctx context.Context, msg string, indent int, args ...any) {
	slog.ErrorContext(ctx, strings.Repeat(" ", indent*int(IndentSpaces))+msg, args...)
}

// InfoIndent calls [slog.Info] on the default logger with indentation.
func InfoIndent(msg string, indent int, args ...any) {
	slog.Info(strings.Repeat(" ", indent*int(IndentSpaces))+msg, args...)
}

// InfoContextIndent calls [slog.InfoContext] on the default logger with indentation.
func InfoContextIndent(ctx context.Context, msg string, indent int, args ...any) {
	slog.InfoContext(ctx, strings.Repeat(" ", indent*int(IndentSpaces))+msg, args...)
}

// trace calls [LevelTrace] on the default logger.
func trace(msg string, args ...any) {
	slog.Default().Log(context.Background(), LevelTrace, msg, args...)
}

// traceContext calls [LevelTrace] on the default logger with context.
func traceContext(ctx context.Context, msg string, args ...any) {
	slog.Default().Log(ctx, LevelTrace, msg, args...)
}
