package model

import (
	"database/sql"
)

func checkResult(r sql.Result) error {
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNoRowsAffected
	}
	return nil
}
