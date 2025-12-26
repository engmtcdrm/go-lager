package examples

import (
	"log/slog"

	"github.com/engmtcdrm/go-lager"
)

// LogTermOnlyExample demonstrates logging only to terminal (stdout and stderr).
func LogTermOnlyExample() {
	logTerm(lager.LevelInfo, false, false)
}

// LogTermDebugOnlyExample demonstrates logging to terminal (stdout and stderr) with debug level.
func LogTermDebugOnlyExample() {
	logTerm(lager.LevelDebug, false, false)
}

// LogTermTraceOnlyExample demonstrates logging to terminal (stdout and stderr) with trace level.
func LogTermTraceOnlyExample() {
	logTerm(lager.LevelTrace, false, false)
}

// LogTermStdoutOnlyExample demonstrates logging only to stdout.
func LogTermStdoutOnlyExample() {
	logTerm(lager.LevelInfo, false, true)
}

// LogTermDebugStdoutOnlyExample demonstrates logging only to stdout with debug level.
func LogTermDebugStdoutOnlyExample() {
	logTerm(lager.LevelDebug, false, true)
}

// LogTermTraceStdoutOnlyExample demonstrates logging only to stdout with trace level.
func LogTermTraceStdoutOnlyExample() {
	logTerm(lager.LevelTrace, false, true)
}

// LogTermStderrOnlyExample demonstrates logging only to stderr.
func LogTermStderrOnlyExample() {
	logTerm(lager.LevelInfo, true, false)
}

// LogTermDebugStderrOnlyExample demonstrates logging only to stderr with debug level.
func LogTermDebugStderrOnlyExample() {
	logTerm(lager.LevelDebug, true, false)
}

// LogTermTraceStderrOnlyExample demonstrates logging only to stderr with trace level.
func LogTermTraceStderrOnlyExample() {
	logTerm(lager.LevelTrace, true, false)
}

// logTerm sets up logging to terminal with specified level and output options.
func logTerm(level slog.Leveler, noStdout, noStderr bool) {
	l := lager.NewLager(level, "").
		NoStdout(noStdout).
		NoStderr(noStderr)

	l.Init()

	doLogging()
}
