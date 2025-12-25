package examples

import (
	"fmt"

	"github.com/engmtcdrm/go-lager"
	pp "github.com/engmtcdrm/go-prettyprint"
)

// LogFileOnlyExample demonstrates logging only to a file.
func LogFileOnlyExample() {
	level := lager.LevelInfo

	logFile := "___app.log"

	fmt.Println("Initializing lager with log file:", pp.Cyan(logFile))
	fmt.Println()

	l := lager.NewLager(level, logFile).
		NoStdout(true).
		NoStderr(true)

	f, err := l.Init()
	if err != nil {
		panic(err)
	}

	if f != nil {
		defer f.Close()
	}

	doLogging()

	readLogContents(logFile)
}
