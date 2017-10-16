package salamander

import (
	"bytes"
	"encoding/json"
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
	return func(s *Server) {
		s.listen = l
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

// binding url path to http handler
func (s *Server) bindRoutes() {
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
		jsonErrResponse(w, http.StatusInternalServerError, err, `encoding json`)
		return
	}
	buf.WriteTo(w)
}

type jsonErr struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func jsonErrResponse(w http.ResponseWriter, st int, err error, msg string) {
	je := jsonErr{
		Message: msg,
	}
	if err != nil {
		je.Error = err.Error()
	}
	jsonResponse(w, st, je)
}
