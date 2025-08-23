package main

import (
	"log/slog"
	"os"

	"github.com/engmtcdrm/go-lager"
)

type Manager struct {
	fileLogger   *slog.Logger
	stdoutLogger *slog.Logger
	stderrLogger *slog.Logger
}

func NewManager(filepath string) *Manager {
	fileLogger := lager.NewFileHandler(filepath)
	stdoutLogger := lager.NewPlainHandler(os.Stdout)
	stderrLogger := lager.NewPlainHandler(os.Stderr)

	return &Manager{
		fileLogger:   slog.New(fileLogger),
		stdoutLogger: slog.New(stdoutLogger),
		stderrLogger: slog.New(stderrLogger),
	}
}

func (m *Manager) Info(msg string, args ...any) {
	m.fileLogger.Info(msg, args...)
	m.stdoutLogger.Info(msg, args...)
	m.stderrLogger.Info(msg, args...)
}

func main() {
	// debug := flag.Bool("debug", false, "enable debug logging")
	// flag.Parse()

	// logFile := lager.Init("___app.log", *debug)
	// defer logFile.Close()

	// indent := 4

	m := NewManager("___app.log")

	m.Info("does this work")

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
	// logger.Info("This is a key check")
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
