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

type Lager struct {
	noFile     bool
	noStdout   bool
	noStderr   bool
	logFileNm  string
	optsFile   *HandlerOptions
	optsStdout *HandlerOptions
	optsStderr *HandlerOptions
}

func NewLager(level slog.Leveler, logFileNm string) *Lager {
	l := &Lager{
		logFileNm: logFileNm,
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

	if logFileNm == "" {
		l.noFile = true
	}

	return l
}

func (l *Lager) NoFile(noFile bool) *Lager {
	l.noFile = noFile
	return l
}

func (l *Lager) NoStdout(noStdout bool) *Lager {
	l.noStdout = noStdout
	return l
}

func (l *Lager) NoStderr(noStderr bool) *Lager {
	l.noStderr = noStderr
	return l
}

func (l *Lager) OptsFile(opts *HandlerOptions) *Lager {
	l.optsFile = opts
	return l
}

func (l *Lager) OptsStdout(opts *HandlerOptions) *Lager {
	l.optsStdout = opts
	return l
}

func (l *Lager) OptsStderr(opts *HandlerOptions) *Lager {
	l.optsStderr = opts
	return l
}

func (l *Lager) Init() (*os.File, error) {
	if l.noFile && l.noStdout && l.noStderr {
		panic("lager: At least one variable, NoFile, NoStdout, or NoStderr must be false")
	}

	var f *os.File
	var err error
	var handlers []slog.Handler

	if !l.noFile {
		f, err = os.OpenFile(l.logFileNm, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
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

func (l *Lager) String() string {
	s := strings.Builder{}
	s.WriteString("Lager Settings:\n")
	s.WriteString(fmt.Sprintf("  Log File Name: %s\n", l.logFileNm))
	s.WriteString(fmt.Sprintf("  Log to File? %t\n", l.noFile))
	s.WriteString(fmt.Sprintf("  Log to Stdout? %t\n", l.noStdout))
	s.WriteString(fmt.Sprintf("  Log to Stderr? %t\n", l.noStderr))
	return s.String()
}
