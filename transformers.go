package lager

import "github.com/engmtcdrm/go-ansi"

// Transform functions for different log levels

// TraceLevelFunc adds magenta color to [LevelTrace] level text in a log statement.
var TraceLevelFunc = func(s string) string {
	return ansi.Magenta + s + ansi.Reset
}

// DebugLevelFunc adds cyan color to [LevelDebug] level text in a log statement.
var DebugLevelFunc = func(s string) string {
	return ansi.Cyan + s + ansi.Reset
}

// InfoLevelFunc leaves [LevelInfo] level text in a log statement unmodified.
var InfoLevelFunc = func(s string) string {
	return s
}

// WarnLevelFunc adds yellow color to [LevelWarning] level text in a log statement.
var WarnLevelFunc = func(s string) string {
	return ansi.Yellow + s + ansi.Reset
}

// ErrorLevelFunc adds red color to [LevelError] level text in a log statement.
var ErrorLevelFunc = func(s string) string {
	return ansi.Red + s + ansi.Reset
}
