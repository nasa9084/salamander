package salamander

// Error is local error type
type Error string

func (e Error) Error() string { return string(e) }

const (
	// ErrDBNotGiven is returned when server is ran without database
	ErrDBNotGiven Error = `database option is not given`

	// ErrScopeRequired in OAuth/OpenID Connect
	ErrScopeRequired Error = `scope parameter is required`
)
