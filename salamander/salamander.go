package salamander

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nasa9084/salamander/salamander/middleware"
)

const (
	defaultListenAddr = ":8080"
)

// Error is local error type
type Error string

func (e Error) Error() string { return string(e) }

const (
	// ErrDBNotGiven is returned when server is ran without database
	ErrDBNotGiven Error = `database option is not given`
)

// server represents a instance of Salamander server
type server struct {
	*mux.Router
	listen string
	mwset  middleware.Set
}

// Server represents a instance of Salamander server
type Server interface {
	Run() error
	Listen() string
}

// NewServer returns a new Salamander server
func NewServer(db *sql.DB, opts ...ServerOption) Server {
	s := server{
		Router: mux.NewRouter(),
		listen: defaultListenAddr,
		mwset: middleware.Set{
			middleware.Transaction(db),
		},
	}
	s.bindRoutes()
	for _, opt := range opts {
		opt(&s)
	}
	return &s
}

// Run server
func (s *server) Run() error {
	return http.ListenAndServe(s.listen, s.mwset.Apply(s.Router))
}

func (s *server) Listen() string {
	return s.listen
}

// binding url path to http handler
func (s *server) bindRoutes() {
	r := s.Router

	r.HandleFunc(`/`, nilHandler)
}
