package loghttp

import (
	"net/http"

	"github.com/aisola/log"
	"github.com/felixge/httpsnoop"
)

var LogRequests = NewRequestLogger(log.DefaultLogger).Middleware

// RequestLogger creates a simple net/http middleware which captures metrics
// about a request/response and then logs it using the the specified Logger in
// aisola/log.
type RequestLogger struct {
	l *log.Logger
}

// NewRequestLogger creates a new request logger from a log.Logger from aisola/log.
func NewRequestLogger(l *log.Logger) *RequestLogger {
	return &RequestLogger{l: l}
}

// Middleware returns the net/http middleware based on this RequestLogger.
func (rl *RequestLogger) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(next, w, r)

		// Log the request
		rl.l.Infof("[%s] %s %s %d %s %d",
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
			m.Code,
			m.Duration,
			m.Written)
	})
}
