package log

import (
	"os"
)

// DefaultLogger is a logger mounted to Stdout.
var DefaultLogger = New(os.Stdout)

// A convenience method mounted to DefaultLogger to Logger.Info.
func Info(message string) {
	DefaultLogger.Info(message)
}

// A convenience method mounted to DefaultLogger to Logger.Infof.
func Infof(format string, args ...interface{}) {
	DefaultLogger.Infof(format, args...)
}

// A convenience method mounted to DefaultLogger to Logger.Debug.
func Debug(message string) {
	DefaultLogger.Debug(message)
}

// A convenience method mounted to DefaultLogger to Logger.Debugf.
func Debugf(format string, args ...interface{}) {
	DefaultLogger.Debugf(format, args...)
}

// A convenience method mounted to DefaultLogger to Logger.SetDebug.
func SetDebug(status bool) {
	DefaultLogger.SetDebug(status)
}
