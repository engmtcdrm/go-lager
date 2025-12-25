// Package lager provides a configurable logging setup using Go's slog package.
// It supports logging to files, standard output, and standard error with
// customizable options for log levels, formatting, and more.
package lager

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

// Lager represents the logging configuration and setup.
type Lager struct {
	noFile     bool
	noStdout   bool
	noStderr   bool
	logFile    string
	optsFile   *HandlerOptions
	optsStdout *HandlerOptions
	optsStderr *HandlerOptions
}

// NewLager creates a new Lager instance with the specified log level and log file name.
func NewLager(level slog.Leveler, logFile string) *Lager {
	l := &Lager{
		logFile: logFile,
		optsFile: &HandlerOptions{
			Level:     level,
			AddTime:   true,
			AddLevel:  true,
			AddSource: true,
		},
		optsStdout: &HandlerOptions{Level: level},
		optsStderr: &HandlerOptions{
			Level:     level,
			AddLevel:  true,
			AddSource: true,
		},
	}

	if logFile == "" {
		l.NoFile(true)
	}

	return l
}

// LogFile sets the log file name.
func (l *Lager) LogFile(logFile string) *Lager {
	l.logFile = logFile
	return l
}

// NoFile sets whether to disable logging to file.
func (l *Lager) NoFile(noFile bool) *Lager {
	l.noFile = noFile
	return l
}

// NoStdout sets whether to disable logging to stdout.
func (l *Lager) NoStdout(noStdout bool) *Lager {
	l.noStdout = noStdout
	return l
}

// NoStderr sets whether to disable logging to stderr.
func (l *Lager) NoStderr(noStderr bool) *Lager {
	l.noStderr = noStderr
	return l
}

// FileOptions sets the handler options for file logging.
func (l *Lager) FileOptions(opts *HandlerOptions) *Lager {
	if opts == nil {
		opts = &HandlerOptions{}
	}

	l.optsFile = opts
	return l
}

// StdoutOptions sets the handler options for stdout logging.
func (l *Lager) StdoutOptions(opts *HandlerOptions) *Lager {
	if opts == nil {
		opts = &HandlerOptions{}
	}

	l.optsStdout = opts
	return l
}

// StderrOptions sets the handler options for stderr logging.
func (l *Lager) StderrOptions(opts *HandlerOptions) *Lager {
	if opts == nil {
		opts = &HandlerOptions{}
	}

	l.optsStderr = opts
	return l
}

// Init initializes the logging system based on the configured settings.
// It returns the file handle if logging to a file is enabled, otherwise nil.
func (l *Lager) Init() (*os.File, error) {
	if l.noFile && l.noStdout && l.noStderr {
		panic("lager: At least one variable, NoFile, NoStdout, or NoStderr must be false")
	}

	var f *os.File
	var err error
	var handlers []slog.Handler

	if !l.noFile {
		f, err = os.OpenFile(l.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return nil, err
		}

		handlerFile := NewFileHandler(f, l.optsFile)
		handlers = append(handlers, handlerFile)
	}

	if !l.noStdout {
		handlerStdout := NewStdoutHandler(l.optsStdout)
		handlers = append(handlers, handlerStdout)
	}

	if !l.noStderr {
		handlerStderr := NewStderrHandler(l.optsStderr)
		handlers = append(handlers, handlerStderr)
	}

	handlerMulti := NewMultiHandler(handlers...)
	logger := slog.New(handlerMulti)
	slog.SetDefault(logger)

	return f, nil
}

// String returns a string representation of the Lager configuration.
func (l *Lager) String() string {
	s := strings.Builder{}
	s.WriteString("Lager Settings:\n")
	s.WriteString(fmt.Sprintf("  Log File Name   : %s\n", l.logFile))
	s.WriteString(fmt.Sprintf("  Log to File     ? %t\n", !l.noFile))
	s.WriteString(fmt.Sprintf("  Log to Stdout   ? %t\n", !l.noStdout))
	s.WriteString(fmt.Sprintf("  Log to Stderr   ? %t\n", !l.noStderr))
	return s.String()
}
