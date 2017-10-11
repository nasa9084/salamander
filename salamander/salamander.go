package salamander

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nasa9084/salamander/salamander/log"
)

// Server represents a instance of Salamander server
type Server struct {
	*mux.Router
	listen       string
	loggingLevel log.Level
}

// NewServer returns a new Salamander server
func NewServer(opts ...option) *Server {
	s := Server{
		Router: mux.NewRouter(),
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
