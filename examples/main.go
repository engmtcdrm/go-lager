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
	slog.Warn(ansi.Yellow + "[Output: FILE | STDERR]: Testing slog warn message" + ansi.Reset)
	slog.Info(ansi.Cyan + "[Output: FILE | STDOUT]: Testing slog info message" + ansi.Reset)
	slog.Debug(ansi.Green + "[Output: FILE | STDERR]: Testing slog debug message" + ansi.Reset)
	lager.Trace(ansi.Magenta + "[Output: FILE | STDERR]: Testing slog trace message" + ansi.Reset)
	fmt.Printf("[Output: %-13s]: Ignore all slogging\n", "STDOUT")
	fmt.Fprintf(os.Stderr, "[Output: %-13s]: Ignore all slogging\n", "STDERR")
}
