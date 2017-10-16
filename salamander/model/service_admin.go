package model

import (
	"database/sql"
	"log"

	"github.com/nasa9084/salamander/salamander/util"
	"github.com/pkg/errors"
)

const (
	serviceAdminCreateSQL = `INSERT INTO service_admin(id, password) VALUES(?, ?)`
	serviceAdminLookupSQL = `SELECT * FROM service_admin WHERE id=?`
	serviceAdminUpdateSQL = `UPDATE service_admin SET(id=?, password=?) WHERE id=?`
	serviceAdminDeleteSQL = `DELETE FROM service_admin WHERE id=?`
)

// ServiceAdmin is a user has authority of System Administrator
type ServiceAdmin struct {
	ID       string `json:"id"`
	Password string `json:"-"`
}

// Scan method
func (sa *ServiceAdmin) Scan(sc scanner) error {
	return sc.Scan(&sa.ID, &sa.Password)
}

// Create ServiceAdmin
func (sa *ServiceAdmin) Create(tx *sql.Tx) error {
	log.Printf("model.ServiceAdmin.Create")

	errmsg := `Creating ServiceAdmin`
	switch {
	case sa.ID == "":
		return errors.Wrap(ErrNilID, errmsg)
	case sa.Password == "":
		return errors.Wrap(ErrNilPasswd, errmsg)
	}

	_, err := tx.Exec(serviceAdminCreateSQL, sa.ID, util.Password(sa.Password, sa.ID))
	if err != nil {
		return errors.Wrap(err, serviceAdminCreateSQL)
	}
	return nil
}

// Lookup ServiceAdmin by ID
func (sa *ServiceAdmin) Lookup(tx *sql.Tx) error {
	log.Printf("model.ServiceAdmin.Lookup")

	if sa.ID == "" {
		return errors.Wrap(ErrNilID, `Looking up ServiceAdmin`)
	}

	row := tx.QueryRow(serviceAdminLookupSQL, sa.ID)
	if err := sa.Scan(row); err != nil {
		return errors.Wrap(err, `Scanning ServiceAdmin`)
	}
	return nil
}

// Update ServiceAdmin information
func (sa *ServiceAdmin) Update(tx *sql.Tx) error {
	log.Printf("model.ServiceAdmin.Update")

	errmsg := `Updating ServiceAdmin`
	switch {
	case sa.ID == "":
		return errors.Wrap(ErrNilID, errmsg)
	case sa.Password == "":
		return errors.Wrap(ErrNilPasswd, errmsg)
	}

	r, err := tx.Exec(serviceAdminUpdateSQL, sa.ID, util.Password(sa.Password, sa.ID), sa.ID)
	if err != nil {
		return errors.Wrap(err, serviceAdminUpdateSQL)
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return errors.Wrap(err, `checking rows affected`)
	}
	if rowsAffected == 0 {
		return ErrNoRowsAffected
	}
	return nil
}

// Delete ServiceAdmin
func (sa *ServiceAdmin) Delete(tx *sql.Tx) error {
	log.Printf("model.ServiceAdmin.Delete")

	if sa.ID == "" {
		return errors.Wrap(ErrNilID, `Deleting ServiceAdmin`)
	}

	r, err := tx.Exec(serviceAdminDeleteSQL, sa.ID)
	if err != nil {
		return errors.Wrap(err, serviceAdminDeleteSQL)
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return errors.Wrap(err, `checking rows affected`)
	}
	if rowsAffected == 0 {
		return ErrNoRowsAffected
	}
	return nil
}
