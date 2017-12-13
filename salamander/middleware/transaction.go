package middleware

import (
	"database/sql"
	"net/http"

	"github.com/nasa9084/salamander/salamander/context"
)

// Transaction embeds database transaction into request.
func Transaction(db *sql.DB) Middleware {
	return &transactionMW{db: db}
}

type transactionMW struct {
	db *sql.DB
}

func (t *transactionMW) Apply(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tx, err := t.db.Begin()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer tx.Rollback()
		r = r.WithContext(context.WithTx(r.Context(), tx))
		h.ServeHTTP(w, r)
	})
}
