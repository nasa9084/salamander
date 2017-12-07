package middleware

import (
	"bytes"
	"net/http"

	"github.com/nasa9084/salamander/salamander/log"
)

const logFormat = `%s %s: %d %s`

type loggingResponseWriter struct {
	w    http.ResponseWriter
	st   int
	body bytes.Buffer
}

func (w *loggingResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.w.Write(b)
}

func (w *loggingResponseWriter) Header() http.Header {
	return w.w.Header()
}

func (w *loggingResponseWriter) WriteHeader(st int) {
	w.st = st
	w.w.WriteHeader(st)
}

// Logger logs all
var Logger Middleware = func(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lw := &loggingResponseWriter{w: w}
		h.ServeHTTP(lw, r)
		log.Info.Printf(logFormat, r.Method, r.URL.Path, lw.st, http.StatusText(lw.st))
	})
}
