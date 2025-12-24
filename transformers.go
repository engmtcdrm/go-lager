package lager

import "github.com/engmtcdrm/go-ansi"

var TraceTransformFunc = func(s string) string {
	return ansi.Magenta + s + ansi.Reset
}

var DebugTransformFunc = func(s string) string {
	return ansi.Cyan + s + ansi.Reset
}

var InfoTransformFunc = func(s string) string {
	return s
}

var WarnTransformFunc = func(s string) string {
	return ansi.Yellow + s + ansi.Reset
}

var ErrorTransformFunc = func(s string) string {
	return ansi.Red + s + ansi.Reset
}
