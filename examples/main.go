package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/engmtcdrm/go-ansi"
	"github.com/engmtcdrm/go-lager"
)

func main() {
	debug := flag.Bool("d", false, "enable debug logging")
	trace := flag.Bool("t", false, "enable trace logging")
	flag.Parse()

	level := lager.LevelInfo

	if *debug {
		level = lager.LevelDebug
	} else if *trace {
		level = lager.LevelTrace
	}

	f, err := lager.Init("___app.log", level)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	slog.Error(ansi.Red + "[Output: FILE | STDERR]: Testing slog error message" + ansi.Reset)
	lager.ErrorIndent(ansi.Red+"[Output: FILE | STDERR]: Testing slog error message"+ansi.Reset, 2)
	slog.Warn(ansi.Yellow + "[Output: FILE | STDERR]: Testing slog warn message" + ansi.Reset)
	lager.WarnIndent(ansi.Yellow+"[Output: FILE | STDERR]: Testing slog warn message"+ansi.Reset, 2)
	slog.Info("[Output: FILE | STDOUT]: Testing slog info message")
	lager.InfoIndent("[Output: FILE | STDOUT]: Testing lager indented info message", 2)
	slog.Debug(ansi.Green + "[Output: FILE | STDERR]: Testing slog debug message" + ansi.Reset)
	lager.DebugIndent(ansi.Green+"[Output: FILE | STDERR]: Testing lager indented debug message"+ansi.Reset, 2)
	lager.Trace(ansi.Magenta + "[Output: FILE | STDERR]: Testing slog trace message" + ansi.Reset)
	lager.TraceIndent(ansi.Magenta+"[Output: FILE | STDERR]: Testing slog trace message"+ansi.Reset, 2)
	fmt.Printf("[Output: %-13s]: Ignore all slogging\n", "STDOUT")
	fmt.Fprintf(os.Stderr, "[Output: %-13s]: Ignore all slogging\n", "STDERR")
}
