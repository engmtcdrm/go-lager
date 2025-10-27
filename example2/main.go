package main

import (
	"log/slog"

	"github.com/engmtcdrm/go-lager"
)

func main() {
	// debug := flag.Bool("debug", false, "enable debug logging")
	// flag.Parse()

	// logFile := lager.Init("___app.log", *debug)
	// defer logFile.Close()

	// indent := 4

	// m := NewManager("___app.log")

	// m.Info("does this work")

	f, err := lager.Init("___app.log", false)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lager.Info("testing slog info")
	lager.Error("testing slog error")
	lager.Debug("testing slog debug")
	lager.Warn("testing slog warn")
	lager.Trace("testing slog trace")

	// lager.Trace("Application started")
	// lager.Debug("This is a debug message")
	// lager.DebugIndent("This is an indented debug message", indent)
	// lager.Debug("")
	// slog.With("indent", indent)
	// ph := lager.NewPlainHandler(os.Stdout)
	// slog.NewTextHandler(os.Stderr, nil)
	// logger := slog.New(ph)
	// logger.WithGroup("main")
	// logger.Info("This is a key check", slog.Int("indent", indent))
	slog.Info("This is a key check")
	// lager.Info("We made it!")
	// lager.InfoIndent("This is indented info", indent)
	// lager.Info("")
	// lager.Warn("This is a warning")
	// lager.WarnIndent("This is an indented warning", indent)
	// lager.Warn("")
	// lager.Error("This has failed badly")
	// lager.ErrorIndent("This is an indented error", indent)
	// lager.Error("")

	// fmt.Print("hellothere!")
	// fmt.Println("only in stdout")
}
