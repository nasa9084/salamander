package salamander

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nasa9084/salamander/salamander/context"
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
	db     *sql.DB
}

// Server represents a instance of Salamander server
type Server interface{
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

// Database configure Database connection
func Database(db *sql.DB) ServerOption {
	return func(s *server) {
		s.db = db
	}
}

// NewServer returns a new Salamander server
func NewServer(opts ...ServerOption) Server {
	s := server{
		Router: mux.NewRouter(),
		listen: defaultListenAddr,
	}
	s.bindRoutes()
	for _, opt := range opts {
		opt(&s)
	}
	return &s
}

// Run server
func (s *server) Run() error {
	if s.db == nil {
		return ErrDBNotGiven
	}
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	return http.ListenAndServe(s.listen, withTx(tx, s.Router))
}

// binding url path to http handler
func (s *server) bindRoutes() {
	r := s.Router

	r.HandleFunc(`/`, nilHandler)
}

func nilHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func jsonResponse(w http.ResponseWriter, st int, v interface{}) {
	w.Header().Set(`Content-Type`, `application/json`)
	w.WriteHeader(st)

	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		jsonResponse(w, http.StatusInternalServerError, newJSONErr(err, `encoding json`))
		return
	}
	buf.WriteTo(w)
}

type jsonErr struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func newJSONErr(err error, msg string) jsonErr {
	je := jsonErr{
		Message: msg,
	}
	if err != nil {
		je.Error = err.Error()
	}
	return je
}

func withTx(tx *sql.Tx, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.WithContext(context.WithTx(r.Context(), tx))
		h.ServeHTTP(w, r)
	})
}
