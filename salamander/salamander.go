package salamander

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	defaultListenAddr = ":8080"
)

// Server represents a instance of Salamander server
type Server struct {
	*mux.Router
	listen string
}


// ServerOption is a functional Option Pattern's option function
type ServerOption func(*Server)

// ListenAddr configure Server listening address
func ListenAddr(l string) ServerOption {
	return func(s *Server)  {
		s.listen = l
		return nil
	}
}

// NewServer returns a new Salamander server
func NewServer(opts ...ServerOption) *Server {
	s := Server{
		Router: mux.NewRouter(),
		listen: ":8080",
	}
	for _, opt := range opts {
		opt(&s)
	}
	s.bindRoutes()
	return &s
}

// Run server
func (s *Server) Run() error {
	return http.ListenAndServe(s.listen, s.Router)
}

func (s *Server) bindRoutes() {
	r := s.Router

	r.HandleFunc(`/`, nilHandler)
}

func nilHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
