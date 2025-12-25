package lager

import (
	"testing"

	"github.com/engmtcdrm/go-ansi"
)

func TestTraceLevelFunc(t *testing.T) {
	input := "TRACE"
	expected := ansi.Magenta + input + ansi.Reset

	result := TraceLevelFunc(input)
	if result != expected {
		t.Errorf("TraceLevelFunc(%q) = %q; want %q", input, result, expected)
	}
}

func TestDebugLevelFunc(t *testing.T) {
	input := "DEBUG"
	expected := ansi.Cyan + input + ansi.Reset

	result := DebugLevelFunc(input)
	if result != expected {
		t.Errorf("DebugLevelFunc(%q) = %q; want %q", input, result, expected)
	}
}

func TestInfoLevelFunc(t *testing.T) {
	input := "INFO"
	expected := input // InfoLevelFunc returns input unmodified

	result := InfoLevelFunc(input)
	if result != expected {
		t.Errorf("InfoLevelFunc(%q) = %q; want %q", input, result, expected)
	}
}

func TestWarnLevelFunc(t *testing.T) {
	input := "WARN"
	expected := ansi.Yellow + input + ansi.Reset

	result := WarnLevelFunc(input)
	if result != expected {
		t.Errorf("WarnLevelFunc(%q) = %q; want %q", input, result, expected)
	}
}

func TestErrorLevelFunc(t *testing.T) {
	input := "ERROR"
	expected := ansi.Red + input + ansi.Reset

	result := ErrorLevelFunc(input)
	if result != expected {
		t.Errorf("ErrorLevelFunc(%q) = %q; want %q", input, result, expected)
	}
}

func TestLevelFuncsWithEmptyString(t *testing.T) {
	tests := []struct {
		name     string
		fn       func(string) string
		expected string
	}{
		{"TraceLevelFunc", TraceLevelFunc, ansi.Magenta + ansi.Reset},
		{"DebugLevelFunc", DebugLevelFunc, ansi.Cyan + ansi.Reset},
		{"InfoLevelFunc", InfoLevelFunc, ""},
		{"WarnLevelFunc", WarnLevelFunc, ansi.Yellow + ansi.Reset},
		{"ErrorLevelFunc", ErrorLevelFunc, ansi.Red + ansi.Reset},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fn("")
			if result != tt.expected {
				t.Errorf("%s(\"\") = %q; want %q", tt.name, result, tt.expected)
			}
		})
	}
}

func TestLevelFuncsWithSpecialCharacters(t *testing.T) {
	input := "TEST\nwith\ttabs and newlines"

	tests := []struct {
		name     string
		fn       func(string) string
		expected string
	}{
		{"TraceLevelFunc", TraceLevelFunc, ansi.Magenta + input + ansi.Reset},
		{"DebugLevelFunc", DebugLevelFunc, ansi.Cyan + input + ansi.Reset},
		{"InfoLevelFunc", InfoLevelFunc, input},
		{"WarnLevelFunc", WarnLevelFunc, ansi.Yellow + input + ansi.Reset},
		{"ErrorLevelFunc", ErrorLevelFunc, ansi.Red + input + ansi.Reset},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fn(input)
			if result != tt.expected {
				t.Errorf("%s(%q) = %q; want %q", tt.name, input, result, tt.expected)
			}
		})
	}
}

func TestLevelFuncsCustomOverwrite(t *testing.T) {
	// Save original functions to restore after test
	originalTrace := TraceLevelFunc
	originalDebug := DebugLevelFunc
	originalInfo := InfoLevelFunc
	originalWarn := WarnLevelFunc
	originalError := ErrorLevelFunc

	// Restore original functions after test
	t.Cleanup(func() {
		TraceLevelFunc = originalTrace
		DebugLevelFunc = originalDebug
		InfoLevelFunc = originalInfo
		WarnLevelFunc = originalWarn
		ErrorLevelFunc = originalError
	})

	// Custom transformer that wraps text in brackets
	bracketWrapper := func(s string) string {
		return "[" + s + "]"
	}

	// Custom transformer that uppercases and adds prefix
	prefixWrapper := func(s string) string {
		return "LOG: " + s
	}

	tests := []struct {
		name       string
		setFunc    func(func(string) string)
		getFunc    func() func(string) string
		customFunc func(string) string
		input      string
		expected   string
	}{
		{
			name:       "TraceLevelFunc custom bracket wrapper",
			setFunc:    func(f func(string) string) { TraceLevelFunc = f },
			getFunc:    func() func(string) string { return TraceLevelFunc },
			customFunc: bracketWrapper,
			input:      "TRACE",
			expected:   "[TRACE]",
		},
		{
			name:       "DebugLevelFunc custom prefix wrapper",
			setFunc:    func(f func(string) string) { DebugLevelFunc = f },
			getFunc:    func() func(string) string { return DebugLevelFunc },
			customFunc: prefixWrapper,
			input:      "DEBUG",
			expected:   "LOG: DEBUG",
		},
		{
			name:       "InfoLevelFunc custom bracket wrapper",
			setFunc:    func(f func(string) string) { InfoLevelFunc = f },
			getFunc:    func() func(string) string { return InfoLevelFunc },
			customFunc: bracketWrapper,
			input:      "INFO",
			expected:   "[INFO]",
		},
		{
			name:       "WarnLevelFunc custom prefix wrapper",
			setFunc:    func(f func(string) string) { WarnLevelFunc = f },
			getFunc:    func() func(string) string { return WarnLevelFunc },
			customFunc: prefixWrapper,
			input:      "WARN",
			expected:   "LOG: WARN",
		},
		{
			name:       "ErrorLevelFunc custom bracket wrapper",
			setFunc:    func(f func(string) string) { ErrorLevelFunc = f },
			getFunc:    func() func(string) string { return ErrorLevelFunc },
			customFunc: bracketWrapper,
			input:      "ERROR",
			expected:   "[ERROR]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set the custom function
			tt.setFunc(tt.customFunc)

			// Verify the function was set correctly
			fn := tt.getFunc()
			result := fn(tt.input)
			if result != tt.expected {
				t.Errorf("custom function returned %q; want %q", result, tt.expected)
			}
		})
	}
}

func TestLevelFuncsOverwriteWithNilSafe(t *testing.T) {
	// Save original function to restore after test
	originalTrace := TraceLevelFunc
	t.Cleanup(func() {
		TraceLevelFunc = originalTrace
	})

	// Test that a no-op function works
	noOpFunc := func(s string) string {
		return ""
	}

	TraceLevelFunc = noOpFunc
	result := TraceLevelFunc("TRACE")
	if result != "" {
		t.Errorf("no-op function returned %q; want empty string", result)
	}
}

func TestLevelFuncsOverwriteChaining(t *testing.T) {
	// Save original function to restore after test
	originalDebug := DebugLevelFunc
	t.Cleanup(func() {
		DebugLevelFunc = originalDebug
	})

	// Test that functions can be chained/composed
	firstWrapper := func(s string) string {
		return "<" + s + ">"
	}
	secondWrapper := func(s string) string {
		return "(" + s + ")"
	}

	// Compose the two wrappers
	DebugLevelFunc = func(s string) string {
		return secondWrapper(firstWrapper(s))
	}

	result := DebugLevelFunc("DEBUG")
	expected := "(<DEBUG>)"
	if result != expected {
		t.Errorf("composed function returned %q; want %q", result, expected)
	}
}
