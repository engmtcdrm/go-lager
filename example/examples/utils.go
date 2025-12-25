package examples

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/engmtcdrm/go-lager"

	pp "github.com/engmtcdrm/go-prettyprint"
)

func doLogging() {
	fmt.Println(pp.Cyan("Logging to various handlers...\n"))
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
	fmt.Print(pp.Cyan("\nFinished logging to various handlers."))
}

func readLogContents(logFile string) {
	fmt.Printf("\n\n")
	fmt.Println(pp.Cyan("Contents of log file:"), pp.Green(logFile))
	fmt.Println()
	rf, _ := os.ReadFile(logFile)
	fmt.Print(string(rf))
	fmt.Print(pp.Cyan("\nEnd of log file contents."))
}
