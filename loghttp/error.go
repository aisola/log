package loghttp

import (
	"net/http"

	"github.com/aisola/log"
)

func Info(w http.ResponseWriter, code int, message string) {
	log.Info(message)
	http.Error(w, message, code)
}

func Error(w http.ResponseWriter, code int, message string, err error) {
	log.Info(message)
	log.Debug(err.Error())
	http.Error(w, message, code)
}
