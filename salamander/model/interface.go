package model

type scanner interface {
	Scan(...interface{}) error
}
