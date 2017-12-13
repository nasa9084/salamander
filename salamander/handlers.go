package salamander

import (
	"bytes"
	"encoding/json"
	"net/http"
)

/*
 * http handling utils.
 */

func nilHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
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
