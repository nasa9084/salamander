package salamander

import "github.com/nasa9084/salamander/salamander/middleware"

// ServerOption is a functional Option Pattern's option function
type ServerOption func(*server)

// ListenAddr configure Server listening address
func ListenAddr(l string) ServerOption {
	return func(s *server) {
		s.listen = l
	}
}

// Middlewares add middlewares to Server.
func Middlewares(mws ...middleware.Middleware) ServerOption {
	return func(s *server) {
		for _, mw := range mws {
			s.mwset = append(s.mwset, mw)
		}
	}
}
