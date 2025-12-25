// Package examples provides a collection of example functions demonstrating various API operations.
package examples

type Example struct {
	Name string
	Fn   func()
}

var AllExamples = []Example{
	{"Log Only to stdout/stderr", LogTermOnlyExample},
	{"Log Only to stdout/stderr with debug level", LogTermDebugOnlyExample},
	{"Log Only to stdout/stderr with trace level", LogTermTraceOnlyExample},
	{"Log Only to stdout", LogTermStdoutOnlyExample},
	{"Log Only to stderr", LogTermStderrOnlyExample},
	{"Log Only to file", LogFileOnlyExample},
}
