package model

import (
	"database/sql"

	"github.com/nasa9084/salamander/salamander/log"
	"github.com/pkg/errors"
)

const (
	corporateCreateSQL = `INSERT INTO corporate(id) VALUES(?)`
	corporateLookupSQL = `SELECT * FROM corporate WHERE id=?`
	corporateUpdateSQL = `UPDATE corporate SET(id=?) WHERE id=?`
	corporateDeleteSQL = `DELETE FROM corporate WHERE id=?`
)

// Corporate model
type Corporate struct {
	ID string `json:"id"`
}

// Scan method
func (c *Corporate) Scan(sc scanner) error {
	return sc.Scan(&c.ID)
}

// Create User
func (c *Corporate) Create(tx *sql.Tx) error {
	log.Info.Printf("model.Corporate.Create")

	errmsg := `Creating Corporate`
	switch {
	case c.ID == "":
		return errors.Wrap(ErrNilID, errmsg)
	}

	_, err := tx.Exec(corporateCreateSQL, c.ID)
	if err != nil {
		return errors.Wrap(err, corporateCreateSQL)
	}
	return nil
}

// Lookup User by ID
func (c *Corporate) Lookup(tx *sql.Tx) error {
	log.Info.Printf("model.Corporate.Lookup")

	if c.ID == "" {
		return errors.Wrap(ErrNilID, `Looking up User`)
	}

	row := tx.QueryRow(corporateLookupSQL, c.ID)
	if err := c.Scan(row); err != nil {
		return errors.Wrap(err, `Scanning User`)
	}
	return nil
}

// Update User
func (c *Corporate) Update(tx *sql.Tx) error {
	log.Info.Printf("model.Corporate.Update")

	errmsg := `Updating User`
	switch {
	case c.ID == "":
		return errors.Wrap(ErrNilID, errmsg)
	}

	r, err := tx.Exec(corporateUpdateSQL, c.ID, c.ID)
	if err != nil {
		return errors.Wrap(err, corporateUpdateSQL)
	}
	return checkResult(r)
}

// Delete User
func (c *Corporate) Delete(tx *sql.Tx) error {
	log.Info.Printf("model.Corporate.Delete")

	if c.ID == "" {
		return errors.Wrap(ErrNilID, `Deleting User`)
	}

	r, err := tx.Exec(corporateDeleteSQL, c.ID)
	if err != nil {
		return errors.Wrap(err, corporateDeleteSQL)
	}
	return checkResult(r)
}
