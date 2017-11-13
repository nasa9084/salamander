package model_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/nasa9084/go-sql-mock"
	"github.com/nasa9084/salamander/salamander/model"
)

var mDB *sql.DB

func TestMain(m *testing.M) {
	setup()
	c := m.Run()
	teardown()
	os.Exit(c)
}

func setup() {
	db, err := sql.Open("sqlmock", "")
	if err != nil {
		os.Exit(1)
	}
	mDB = db
}

func teardown() {
	mDB.Close()
}

type mock struct{}

func (m mock) GetCreateSQL() string           { return "" }
func (m mock) GetCreateValues() []interface{} { return []interface{}{} }

func TestCreate(t *testing.T) {
	candidates := []struct {
		Label          string
		Model          model.CreateModel
		ExpectedResult sqlmock.ResultOpts
		Expected       error
	}{
		{"empty result", mock{}, sqlmock.RowsAffected(0), model.ErrNoRowsAffected},
		{"valid", mock{}, sqlmock.RowsAffected(1), nil},
		{"nil value", nil, sqlmock.RowsAffected(1), model.ErrNilModel},
	}
	for _, c := range candidates {
		t.Log(c.Label)
		sqlmock.ExpectedResult(c.ExpectedResult)
		tx, err := mDB.Begin()
		if err != nil {
			t.Fatal(err)
		}
		if err = model.Create(tx, c.Model); err != c.Expected {
			t.Errorf("error should be %s, but actual %s", c.Expected, err)
			return
		}
	}
}

func TestRead(t *testing.T) {
}

func TestUpdate(t *testing.T) {
}

func TestDelete(t *testing.T) {
}
