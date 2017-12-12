package salamander

import (
	"bytes"
	"encoding/json"
	"net/http"
)

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
