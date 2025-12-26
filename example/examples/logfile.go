package examples

import (
	"fmt"
	"log/slog"

	"github.com/engmtcdrm/go-lager"
	pp "github.com/engmtcdrm/go-prettyprint"
)

// LogFileOnlyExample demonstrates logging only to a file.
func LogFileOnlyExample() {
	logFile(lager.LevelInfo, false, false)
}

// LogFileDebugOnlyExample demonstrates logging only to a file with debug level.
func LogFileDebugOnlyExample() {
	logFile(lager.LevelDebug, false, false)
}

// LogFileTraceOnlyExample demonstrates logging only to a file with trace level.
func LogFileTraceOnlyExample() {
	logFile(lager.LevelTrace, false, false)
}

// LogFileStdoutOnlyExample demonstrates logging only to stdout.
func LogFileStdoutOnlyExample() {
	logFile(lager.LevelInfo, false, true)
}

// LogFileDebugStdoutOnlyExample demonstrates logging only to stdout with debug level.
func LogFileDebugStdoutOnlyExample() {
	logFile(lager.LevelDebug, false, true)
}

// LogFileTraceStdoutOnlyExample demonstrates logging only to stdout with trace level.
func LogFileTraceStdoutOnlyExample() {
	logFile(lager.LevelTrace, false, true)
}

// LogFileStderrOnlyExample demonstrates logging only to stderr.
func LogFileStderrOnlyExample() {
	logFile(lager.LevelInfo, true, false)
}

// LogFileDebugStderrOnlyExample demonstrates logging only to stderr with debug level.
func LogFileDebugStderrOnlyExample() {
	logFile(lager.LevelDebug, true, false)
}

// LogFileTraceStderrOnlyExample demonstrates logging only to stderr with trace level.
func LogFileTraceStderrOnlyExample() {
	logFile(lager.LevelTrace, true, false)
}

// logFile sets up logging to a file with specified level.
func logFile(level slog.Leveler, noStdout, noStderr bool) {
	logFile := "___app.log"

	fmt.Println("Initializing lager with log file:", pp.Cyan(logFile))
	fmt.Println()

	l := lager.NewLager(level, logFile).
		NoStdout(noStdout).
		NoStderr(noStderr)

	f, err := l.Init()
	if err != nil {
		panic(err)
	}

	if f != nil {
		defer f.Close()
	}

	doLogging()

	readLogContents(logFile)
}
