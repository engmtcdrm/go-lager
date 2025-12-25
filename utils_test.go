package lager

import (
	"fmt"
	"log/slog"
	"testing"
)

func TestLevelString(t *testing.T) {
	tests := []struct {
		level    slog.Level
		expected string
	}{
		{slog.Level(-11), "TRACE-3"},
		{slog.Level(-10), "TRACE-2"},
		{slog.Level(-9), "TRACE-1"},
		{LevelTrace, "TRACE"},
		{slog.Level(-7), "TRACE+1"},
		{slog.Level(-6), "TRACE+2"},
		{slog.Level(-5), "TRACE+3"},
		{LevelDebug, "DEBUG"},
		{slog.Level(-3), "DEBUG+1"},
		{slog.Level(-2), "DEBUG+2"},
		{slog.Level(-1), "DEBUG+3"},
		{LevelInfo, "INFO"},
		{slog.Level(1), "INFO+1"},
		{slog.Level(2), "INFO+2"},
		{slog.Level(3), "INFO+3"},
		{LevelWarn, "WARN"},
		{slog.Level(5), "WARN+1"},
		{slog.Level(6), "WARN+2"},
		{slog.Level(7), "WARN+3"},
		{LevelError, "ERROR"},
		{slog.Level(9), "ERROR+1"},
		{slog.Level(10), "ERROR+2"},
		{slog.Level(11), "ERROR+3"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("level_%d", tt.level), func(t *testing.T) {
			result := levelString(tt.level)
			if result != tt.expected {
				t.Errorf("levelString(%d) = %q; want %q", tt.level, result, tt.expected)
			}
		})
	}
}
