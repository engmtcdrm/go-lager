package main

import (
	"flag"
	"fmt"

	"github.com/engmtcdrm/go-lager"
)

func main() {
	debug := flag.Bool("debug", false, "enable debug logging")
	flag.Parse()

	logFile := lager.Init("___app.log", *debug)
	defer logFile.Close()

	indent := 4

	lager.Trace("Application started")
	lager.Debug("This is a debug message")
	lager.DebugIndent("This is an indented debug message", indent)
	lager.Debug("")
	lager.Info("We made it!")
	lager.InfoIndent("This is indented info", indent)
	lager.Info("")
	lager.Warn("This is a warning")
	lager.WarnIndent("This is an indented warning", indent)
	lager.Warn("")
	lager.Error("This has failed badly")
	lager.ErrorIndent("This is an indented error", indent)
	lager.Error("")

	fmt.Print("hellothere!")
	fmt.Println("only in stdout")
}
