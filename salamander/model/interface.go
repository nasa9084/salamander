package model

type scanner interface {
	Scan(...interface{}) error
}

// CRUDModel implements CreateModel, ReadModel, UpdateModel, and DeleteModel.
type CRUDModel interface {
	CreateModel
	ReadModel
	UpdateModel
	DeleteModel
}

// CreateModel can create a new database record.
type CreateModel interface {
	GetCreateSQL() string
	GetCreateValues() []interface{}
}

// ReadModel can read(lookup) a record from database.
type ReadModel interface {
	GetReadSQL() string
	GetReadValues() []interface{}
	Scan(scanner) error
}

// UpdateModel can update a database record.
type UpdateModel interface {
	GetUpdateSQL() string
	GetUpdateValues() []interface{}
}

// DeleteModel can delete a database record.
type DeleteModel interface {
	GetDeleteSQL() string
	GetDeleteValues() []interface{}
}
