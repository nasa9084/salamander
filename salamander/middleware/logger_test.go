package middleware_test

import (
	"bytes"
	"log"
	"net/http"
	"strings"
	"testing"

	logger "github.com/nasa9084/salamander/salamander/log"
	"github.com/nasa9084/salamander/salamander/middleware"
)

func TestLogger(t *testing.T) {
	stdout := bytes.Buffer{}
	logger.Info = log.New(&stdout, "", log.Ldate|log.Ltime)

	candidates := []struct {
		method   string
		url      string
		expected string
	}{
		{"GET", "http://localhost", "GET /: 200 OK"},
		{"GET", "http://localhost/", "GET /: 200 OK"},
		{"GET", "http://localhost/hoge", "GET /hoge: 200 OK"},
		{"GET", "http://localhost?foo=bar", "GET /: 200 OK"},
		{"GET", "http://localhost/hoge?foo=bar", "GET /hoge: 200 OK"},
	}
	for _, c := range candidates {
		stdout.Reset()
		w := newResponseWriter()
		r, err := http.NewRequest(c.method, c.url, &bytes.Buffer{})
		if err != nil {
			t.Fatal(err)
		}
		h := middleware.Logger().Apply(http.HandlerFunc(nilHandler))
		h.ServeHTTP(w, r)
		if !strings.HasSuffix(strings.TrimSpace(stdout.String()), c.expected) {
			t.Errorf("%s != %s", strings.TrimSpace(stdout.String()), c.expected)
			return
		}
	}
}
