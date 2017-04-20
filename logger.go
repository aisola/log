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

// Writes a basic debug-level message to the logger, then a newline. The message
// will only write if you have called Logger.SetDebug with true.
func (l *Logger) Debug(message string) {
	if l.debug {
		fmt.Fprintf(l.w, "[DEBUG] %s\n", message)
	}
}

// Writes a formatted debug-level message to the logger, then a newline. The
// message will only write if you have called Logger.SetDebug with true.
func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.debug {
		format = fmt.Sprintf("[DEBUG] %s\n", format)
		fmt.Fprintf(l.w, format, args...)
	}
}

// Write a basic info-level message to the logger, then a newline.
func (l *Logger) Info(message string) {
	fmt.Fprintf(l.w, "[INFO ] %s\n", message)
}

// Write a formatted info-level message to the logger, then a newline.
func (l *Logger) Infof(format string, args ...interface{}) {
	format = fmt.Sprintf("[INFO ] %s\n", format)
	fmt.Fprintf(l.w, format, args...)
}

// Sets the log level. If called with true, then the logger will write the
// debug-level log messages.
func (l *Logger) SetDebug(status bool) {
	l.debug = status
}
