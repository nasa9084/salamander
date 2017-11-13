package model

import (
	"github.com/nasa9084/salamander/salamander/util"
)

const (
	userCreateSQL = `INSERT INTO user(id, password, display_name, email) VALUES(?, ?, ?, ?)`
	userLookupSQL = `SELECT * FROM user WHERE id=?`
	userUpdateSQL = `UPDATE user SET(id=?, password=?, display_name=?, email=?) WHERE id=?`
	userDeleteSQL = `DELETE FROM user WHERE id=?`
)

// User is a user, such as corporate admin, arbeit
type User struct {
	ID       string `json:"id"`
	Password string `json:"-"`

	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

// Scan method
func (u *User) Scan(sc scanner) error {
	return sc.Scan(&u.ID, &u.Password, &u.DisplayName, &u.Email)
}

// GetCreateSQL returns SQL query for creating a new database record.
func (u *User) GetCreateSQL() string { return userCreateSQL }

// GetCreateValues returns values list for placeholders in query returned GetCreateSQL().
func (u *User) GetCreateValues() []interface{} {
	return []interface{}{u.ID, util.Password(u.Password, u.ID), u.DisplayName, u.Email}
}

// GetReadSQL returns SQL query for read from database record.
func (u *User) GetReadSQL() string { return userLookupSQL }

// GetReadValues returns values list for placeholders in query returned GetReadSQL().
func (u *User) GetReadValues() []interface{} { return []interface{}{u.ID} }

// GetUpdateSQL returns SQL query for update a database record.
func (u *User) GetUpdateSQL() string { return userUpdateSQL }

// GetUpdateValues returns values list for placeholders in query returned GetUpdateSQL().
func (u *User) GetUpdateValues() []interface{} {
	return []interface{}{u.ID, util.Password(u.Password, u.ID), u.DisplayName, u.Email, u.ID}
}

// GetDeleteSQL returns SQL query for delete a database record.
func (u *User) GetDeleteSQL() string { return userDeleteSQL }

// GetDeleteValues returns values list for placeholders in query returned GetDeleteSQL().
func (u *User) GetDeleteValues() []interface{} { return []interface{}{u.ID} }
