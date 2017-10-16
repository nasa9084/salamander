package model

import (
	"database/sql"

	"github.com/pkg/errors"
)

func checkResult(r sql.Result) error {
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return errors.Wrap(err, `checking rows affected`)
	}
	if rowsAffected == 0 {
		return ErrNoRowsAffected
	}
	return nil
}
