package loghttp

import (
	"net/http"

	"github.com/aisola/log"
)

// Info is a convenience function which will pass the given message to log.Info
// then call http.Error with the code and message.
func Info(w http.ResponseWriter, code int, message string) {
	log.Info(message)
	http.Error(w, message, code)
}

// Error is a convenience function which will pass the givem message to log.Info,
// the given error to log.Debug, and then calls http.Error with the message and
// given code.
func Error(w http.ResponseWriter, code int, message string, err error) {
	log.Info(message)
	log.Debug(err.Error())
	http.Error(w, message, code)
}
