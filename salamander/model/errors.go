package model

import "errors"

var (
	// ErrNilID is returned when the object's ID is nil
	ErrNilID = errors.New(`object's ID is nil`)

	// ErrNilPasswd is returned when creating or updating but the object's Password is nil
	ErrNilPasswd = errors.New(`object's Password is nil`)
)
