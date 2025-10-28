package main

import (
	"flag"
	"fmt"
	"log/slog"

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

	slog.Info("(FILE|STDOUT): testing slog info")
	slog.Error("(FILE|STDERR): testing slog error")
	slog.Debug("(FILE|STDERR): testing slog debug")
	slog.Warn("(FILE|STDERR): testing slog warn")
	lager.Trace("(FILE|STDERR): testing slog trace")
	fmt.Println("(STDOUT): hello there!")
}
