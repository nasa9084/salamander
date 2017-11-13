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

func validateValues(values []interface{}) error {
	for _, val := range values {
		if val == nil {
			return ErrNilValue
		}
	}
	return nil
}

func exec(tx *sql.Tx, q string, v []interface{}) error {
	if err := validateValues(v); err != nil {
		return err
	}
	stmt, err := tx.Prepare(q)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(v...)
	if err != nil {
		return err
	}
	return checkResult(res)
}

func queryRow(tx *sql.Tx, q string, v []interface{}) (*sql.Row, error) {
	if err := validateValues(v); err != nil {
		return nil, err
	}
	stmt, err := tx.Prepare(q)
	if err != nil {
		return nil, err
	}
	return stmt.QueryRow(v...), nil
}

// Create a new database record from model.
func Create(tx *sql.Tx, m CreateModel) error {
	if m == nil {
		return ErrNilModel
	}
	q := m.GetCreateSQL()
	v := m.GetCreateValues()
	return exec(tx, q, v)
}

// Read from database.
func Read(tx *sql.Tx, m ReadModel) error {
	q := m.GetReadSQL()
	v := m.GetReadValues()
	row, err := queryRow(tx, q, v)
	if err != nil {
		return err
	}
	return m.Scan(row)
}

// Update a database record.
func Update(tx *sql.Tx, m UpdateModel) error {
	q := m.GetUpdateSQL()
	v := m.GetUpdateValues()
	return exec(tx, q, v)
}

// Delete a database record.
func Delete(tx *sql.Tx, m DeleteModel) error {
	q := m.GetDeleteSQL()
	v := m.GetDeleteValues()
	return exec(tx, q, v)
}
