package salamander

import (
	"bytes"
	"database/sql"
	"encoding/json"
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
}

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

// binding url path to http handler
func (s *server) bindRoutes() {
	r := s.Router

	r.HandleFunc(`/`, nilHandler)
}

func nilHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

func jsonResponse(w http.ResponseWriter, st int, v interface{}) {
	w.Header().Set(`Content-Type`, `application/json`)
	w.WriteHeader(st)

	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		jsonError(w, http.StatusInternalServerError, err, `encoding json`)
		return
	}
	buf.WriteTo(w)
}

type jsonErr struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func jsonError(w http.ResponseWriter, st int, err error, msg string) {
	je := jsonErr{
		Message: msg,
	}
	if err != nil {
		je.Error = err.Error()
	}
	jsonResponse(w, st, je)
}
