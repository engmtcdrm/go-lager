package lager

import (
	"fmt"
	"log/slog"
)

func levelString(l slog.Level) string {
	if l < slog.LevelDebug {
		str := func(base string, val slog.Level) string {
			if val == 0 {
				return base
			}
			return fmt.Sprintf("%s%+d", base, val)
		}
		return str("TRACE", l-LevelTrace)
	}

	return l.String()
}
