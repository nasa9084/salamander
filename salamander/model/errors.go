package model

// Error is package specific error type
type Error string

const (
	// ErrNoRowsAffected is returned if exec query affect to no rows.
	ErrNoRowsAffected Error = "no rows affected"
	// ErrNilModel is returned if given model is nil.
	ErrNilModel Error = "given model is nil"
	// ErrNilValue is returned if values validation got error.
	ErrNilValue Error = "one or some value is nil"
)

func (err Error) Error() string { return string(err) }
