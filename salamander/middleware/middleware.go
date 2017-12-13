package middleware

import "net/http"

// Middleware is a HTTP handler middleware.
type Middleware interface {
	Apply(http.Handler) http.Handler
}

// Set is a slice of Middleware.
type Set []Middleware

// Apply all of listed middlewares.
func (mwset Set) Apply(h http.Handler) http.Handler {
	for _, mw := range mwset {
		h = mw.Apply(h)
	}
	return h
}
