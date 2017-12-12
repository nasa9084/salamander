package middleware_test

import (
	"net/http"
	"os"
	"testing"
)

type responseWriter struct {
	header http.Header
	status int
	body   []byte
}

func newResponseWriter() http.ResponseWriter {
	return &responseWriter{
		header: http.Header{},
	}
}

func (w *responseWriter) Header() http.Header {
	return w.header
}

func (w *responseWriter) Write(b []byte) (int, error) {
	return copy(w.body, b), nil
}

func (w *responseWriter) WriteHeader(st int) {
	w.status = st
}

func nilHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
