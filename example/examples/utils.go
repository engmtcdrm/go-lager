package examples

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/engmtcdrm/go-lager"

	pp "github.com/engmtcdrm/go-prettyprint"
)

// doLogging performs various logging operations to demonstrate the configured handlers.
func doLogging() {
	cFile := pp.Cyan("FILE")
	cStdout := pp.Green("STDOUT")
	cStderr := pp.Red("STDERR")

	fmt.Println(pp.Cyan("Logging to various handlers...\n"))
	slog.Error(fmt.Sprintf("[Output: %s | %s]: Testing slog error message using [%s]", cFile, cStderr, pp.Yellow("slog.Error")))
	lager.ErrorIndent(fmt.Sprintf("[Output: %s | %s]: Testing slog error message using [%s]", cFile, cStderr, pp.Yellow("lager.ErrorIndent")), 1)
	slog.Warn(fmt.Sprintf("[Output: %s | %s]: Testing slog warn message using [%s]", cFile, cStderr, pp.Yellow("slog.Warn")))
	lager.WarnIndent(fmt.Sprintf("[Output: %s | %s]: Testing slog warn message using [%s]", cFile, cStderr, pp.Yellow("lager.WarnIndent")), 1)
	slog.Info(fmt.Sprintf("[Output: %s | %s]: Testing slog info message using [%s]", cFile, cStdout, pp.Yellow("slog.Info")))
	lager.InfoIndent(fmt.Sprintf("[Output: %s | %s]: Testing lager indented info message using [%s]", cFile, cStdout, pp.Yellow("lager.InfoIndent")), 1)
	slog.Debug(fmt.Sprintf("[Output: %s | %s]: Testing slog debug message using [%s]", cFile, cStderr, pp.Yellow("slog.Debug")))
	lager.DebugIndent(fmt.Sprintf("[Output: %s | %s]: Testing lager indented debug message using [%s]", cFile, cStderr, pp.Yellow("lager.DebugIndent")), 1)
	lager.Trace(fmt.Sprintf("[Output: %s | %s]: Testing slog trace message using [%s]", cFile, cStderr, pp.Yellow("lager.Trace")))
	lager.TraceIndent(fmt.Sprintf("[Output: %s | %s]: Testing slog trace message using [%s]", cFile, cStderr, pp.Yellow("lager.TraceIndent")), 1)
	fmt.Printf("[Output: %s]: Ignore all slogging using [%s]\n", cStdout, pp.Yellow("fmt.Printf"))
	fmt.Fprintf(os.Stderr, "[Output: %s]: Ignore all slogging using [%s]\n", cStderr, pp.Yellow("fmt.Fprintf"))
	fmt.Print(pp.Cyan("\nFinished logging to various handlers."))
}

// readLogContents reads and prints the contents of the specified log file.
func readLogContents(logFile string) {
	fmt.Printf("\n\n")
	fmt.Println(pp.Cyan("Contents of log file:"), pp.Green(logFile))
	fmt.Println()
	rf, _ := os.ReadFile(logFile)
	fmt.Print(string(rf))
	fmt.Print(pp.Cyan("\nEnd of log file contents."))
}
