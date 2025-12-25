package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/engmtcdrm/go-ansi"
	"github.com/engmtcdrm/go-lager"
)

func logIt() {
	slog.Error("[Output: FILE | STDERR]: Testing slog error message")
	lager.ErrorIndent("[Output: FILE | STDERR]: Testing slog error message", 1)
	slog.Warn("[Output: FILE | STDERR]: Testing slog warn message")
	lager.WarnIndent("[Output: FILE | STDERR]: Testing slog warn message", 1)
	slog.Info("[Output: FILE | STDOUT]: Testing slog info message")
	lager.InfoIndent("[Output: FILE | STDOUT]: Testing lager indented info message", 1)
	slog.Debug("[Output: FILE | STDERR]: Testing slog debug message")
	lager.DebugIndent("[Output: FILE | STDERR]: Testing lager indented debug message", 1)
	lager.Trace("[Output: FILE | STDERR]: Testing slog trace message")
	lager.TraceIndent("[Output: FILE | STDERR]: Testing slog trace message", 1)
	fmt.Printf("[Output: %-13s]: Ignore all slogging\n", "STDOUT")
	fmt.Fprintf(os.Stderr, "[Output: %-13s]: Ignore all slogging\n", "STDERR")
}

func cyanMe(s string) string {
	return ansi.Cyan + ansi.Bold + s + ansi.Reset
}

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

	logFile := "___app.log"

	l := lager.NewLager(level, logFile).
		NoFile(false).
		NoStderr(false).
		NoStdout(false)

	fmt.Println("Initializing lager with log file:", cyanMe(logFile))
	fmt.Println(l)

	f, err := l.Init()
	if err != nil {
		panic(err)
	}

	if f != nil {
		defer f.Close()
	}

	fmt.Println(cyanMe("Logging to various handlers...\n"))

	logIt()

	fmt.Println()
	fmt.Printf(cyanMe("Contents of log file: %s\n\n"), cyanMe(logFile))

	rf, _ := os.ReadFile(logFile)
	fmt.Print(string(rf))

	fmt.Println()
	fmt.Println(cyanMe("Updating lager to not log to file...\n"))
	l.NoFile(true)
	l.Init()
	fmt.Println(l)

	logIt()

	fmt.Println()
	fmt.Printf(cyanMe("Contents of log file should be the same as before: %s\n\n"), ansi.Yellow+logFile+ansi.Reset)

	rf, _ = os.ReadFile(logFile)
	fmt.Print(string(rf))

	fmt.Println()
	fmt.Println(cyanMe(fmt.Sprintf("Updating lager to not log to %s...\n", ansi.Yellow+"stdout"+ansi.Cyan)))
	l.NoStdout(true)
	l.Init()
	fmt.Println(l)

	logIt()

	fmt.Println()
	fmt.Println(cyanMe(fmt.Sprintf("Updating lager to not log to %s...\n", ansi.Yellow+"stderr"+ansi.Cyan)))
	l.NoStdout(false).NoStderr(true)
	l.Init()
	fmt.Println(l)

	logIt()
}
