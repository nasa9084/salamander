package model

import (
	"database/sql"

	"github.com/nasa9084/salamander/salamander/log"
	"github.com/nasa9084/salamander/salamander/util"
	"github.com/pkg/errors"
)

const (
	userCreateSQL = `INSERT INTO user(id, password, first_name, first_name_kana, family_name, family_name_kana) VALUES(?, ?, ?, ?, ?, ?)`
	userLookupSQL = `SELECT * FROM user WHERE id=?`
	userUpdateSQL = `UPDATE user SET(id=?, password=?, first_name=?, first_name_kana=?, family_name=?, family_name_kana=?) WHERE id=?`
	userDeleteSQL = `DELETE FROM user WHERE id=?`
)

// User is a user, such as corporate admin, arbeit
type User struct {
	ID             string `json:"id"`
	Password       string `json:"-"`
	FirstName      string `json:"first_name"`
	FirstNameKana  string `json:"first_name_kana"`
	FamilyName     string `json:"family_name"`
	FamilyNameKana string `json:"family_name_kana"`
}

// Scan method
func (u *User) Scan(sc scanner) error {
	return sc.Scan(&u.ID, &u.Password, &u.FirstName, &u.FirstNameKana, &u.FamilyName, &u.FamilyNameKana)
}

// Create User
func (u *User) Create(tx *sql.Tx) error {
	log.Info.Printf("model.User.Create")

	errmsg := `Creating User`
	switch {
	case u.ID == "":
		return errors.Wrap(ErrNilID, errmsg)
	case u.Password == "":
		return errors.Wrap(ErrNilPasswd, errmsg)
	}

	_, err := tx.Exec(userCreateSQL, u.ID, util.Password(u.Password, u.ID), u.FirstName, u.FirstNameKana, u.FamilyName, u.FamilyNameKana)
	if err != nil {
		return errors.Wrap(err, userCreateSQL)
	}
	return nil
}

// Lookup User by ID
func (u *User) Lookup(tx *sql.Tx) error {
	log.Info.Printf("model.User.Lookup")

	if u.ID == "" {
		return errors.Wrap(ErrNilID, `Looking up User`)
	}

	row := tx.QueryRow(userLookupSQL, u.ID)
	if err := u.Scan(row); err != nil {
		return errors.Wrap(err, `Scanning User`)
	}
	return nil
}

// Update User
func (u *User) Update(tx *sql.Tx) error {
	log.Info.Printf("model.User.Update")

	errmsg := `Updating User`
	switch {
	case u.ID == "":
		return errors.Wrap(ErrNilID, errmsg)
	case u.Password == "":
		return errors.Wrap(ErrNilPasswd, errmsg)
	}

	r, err := tx.Exec(userUpdateSQL, u.ID, util.Password(u.Password, u.ID), u.FirstName, u.FirstNameKana, u.FamilyName, u.FamilyNameKana, u.ID)
	if err != nil {
		return errors.Wrap(err, userUpdateSQL)
	}
	return checkResult(r)
}

// Delete User
func (u *User) Delete(tx *sql.Tx) error {
	log.Info.Printf("model.User.Delete")

	if u.ID == "" {
		return errors.Wrap(ErrNilID, `Deleting User`)
	}

	r, err := tx.Exec(userDeleteSQL, u.ID)
	if err != nil {
		return errors.Wrap(err, userDeleteSQL)
	}
	return checkResult(r)
}
