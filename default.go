package log

import (
	"os"
)

// DefaultLogger is a logger mounted to Stdout.
var DefaultLogger = New(os.Stdout)

// Info is a convenience method mounted to DefaultLogger to Logger.Info.
func Info(message string) {
	DefaultLogger.Info(message)
}

// Infof is a convenience method mounted to DefaultLogger to Logger.Infof.
func Infof(format string, args ...interface{}) {
	DefaultLogger.Infof(format, args...)
}

// Debug is a convenience method mounted to DefaultLogger to Logger.Debug.
func Debug(message string) {
	DefaultLogger.Debug(message)
}

// Debugf is a convenience method mounted to DefaultLogger to Logger.Debugf.
func Debugf(format string, args ...interface{}) {
	DefaultLogger.Debugf(format, args...)
}

// SetDebug is a convenience method mounted to DefaultLogger to Logger.SetDebug.
func SetDebug(status bool) {
	DefaultLogger.SetDebug(status)
}
