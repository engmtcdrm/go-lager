// Package examples provides a collection of example functions demonstrating various API operations.
package examples

import (
	"github.com/engmtcdrm/go-eggy"
	"github.com/engmtcdrm/go-lager"
)

var AllExamples = []eggy.Example{
	{Name: "Log to file and stdout/stderr", Fn: LogAllExample},
	{Name: "Log to file and stdout/stderr with debug level", Fn: LogAllDebugExample},
	{Name: "Log to file and stdout/stderr with trace level", Fn: LogAllTraceExample},
	{Name: "Log Only to stdout/stderr", Fn: LogTermOnlyExample},
	{Name: "Log Only to stdout/stderr with debug level", Fn: LogTermDebugOnlyExample},
	{Name: "Log Only to stdout/stderr with trace level", Fn: LogTermTraceOnlyExample},
	{Name: "Log Only to stdout", Fn: LogTermStdoutOnlyExample},
	{Name: "Log Only to stdout with debug level", Fn: LogTermDebugStdoutOnlyExample},
	{Name: "Log Only to stdout with trace level", Fn: LogTermTraceStdoutOnlyExample},
	{Name: "Log Only to stderr", Fn: LogTermStderrOnlyExample},
	{Name: "Log Only to stderr with debug level", Fn: LogTermDebugStderrOnlyExample},
	{Name: "Log Only to stderr with trace level", Fn: LogTermTraceStderrOnlyExample},
	{Name: "Log Only to file", Fn: LogFileOnlyExample},
	{Name: "Log Only to file with debug level", Fn: LogFileDebugOnlyExample},
	{Name: "Log Only to file with trace level", Fn: LogFileTraceOnlyExample},
	{Name: "Log Only to file and stdout", Fn: LogFileStdoutOnlyExample},
	{Name: "Log Only to file and stdout with debug level", Fn: LogFileDebugStdoutOnlyExample},
	{Name: "Log Only to file and stdout with trace level", Fn: LogFileTraceStdoutOnlyExample},
	{Name: "Log Only to file and stderr", Fn: LogFileStderrOnlyExample},
	{Name: "Log Only to file and stderr with debug level", Fn: LogFileDebugStderrOnlyExample},
	{Name: "Log Only to file and stderr with trace level", Fn: LogFileTraceStderrOnlyExample},
}

// LogAllExample demonstrates logging to file and terminal (stdout and stderr).
func LogAllExample() {
	logFile(lager.LevelInfo, false, false)
}

// LogAllDebugExample demonstrates logging to file and terminal (stdout and stderr) with debug level.
func LogAllDebugExample() {
	logFile(lager.LevelDebug, false, false)
}

// LogAllTraceExample demonstrates logging to file and terminal (stdout and stderr) with trace level.
func LogAllTraceExample() {
	logFile(lager.LevelTrace, false, false)
}
