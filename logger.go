// Package log intends to be an extremely simple, lightweight logging
// library with only what is necessary for decent logging in most
// golang programs
package log

import (
	"fmt"
	"io"
)

// Logger is the structure which makes up a place to log.
type Logger struct {
	debug bool
	w     io.Writer
}

// New will take an io.Writer and return a logger.
func New(w io.Writer) *Logger {
	return &Logger{
		debug: false,
		w:     w,
	}
}

// NewMultiLogger is a convenience function to log to multiple sources. This is
// equivalent to calling New and passing in the result from io.MultiWriter.
func NewMultiLogger(w ...io.Writer) *Logger {
	return New(io.MultiWriter(w...))
}

// Debug writes a basic debug-level message to the logger, then a newline. The
// message will only write if you have called Logger.SetDebug with true.
func (l *Logger) Debug(message string) {
	if l.debug {
		fmt.Fprintf(l.w, "[DEBUG] %s\n", message)
	}
}

// Debugf writes a formatted debug-level message to the logger, then a newline.
// The message will only write if you have called Logger.SetDebug with true.
func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.debug {
		format = fmt.Sprintf("[DEBUG] %s\n", format)
		fmt.Fprintf(l.w, format, args...)
	}
}

// Info writes a basic info-level message to the logger, then a newline.
func (l *Logger) Info(message string) {
	fmt.Fprintf(l.w, "[INFO ] %s\n", message)
}

// Infof writes a formatted info-level message to the logger, then a newline.
func (l *Logger) Infof(format string, args ...interface{}) {
	format = fmt.Sprintf("[INFO ] %s\n", format)
	fmt.Fprintf(l.w, format, args...)
}

// SetDebug sets the log level to write debug messages when called with true,
// otherwise sets the log level to not write debug messages.
func (l *Logger) SetDebug(status bool) {
	l.debug = status
}
