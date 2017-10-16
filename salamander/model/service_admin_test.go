package model_test

import (
	"database/sql"
	"database/sql/driver"
	"testing"

	sqlmock "github.com/nasa9084/go-sql-mock"
	"github.com/nasa9084/salamander/salamander/model"
	"github.com/pkg/errors"
)

func transaction(t *testing.T) *sql.Tx {
	tx, err := mockDB.Begin()
	if err != nil {
		t.Fatalf("%s", err)
	}
	return tx
}

func TestServiceAdminCreate(t *testing.T) {
	candidates := []struct {
		Name        string
		ID          string
		Password    string
		ExpectedErr error
	}{
		{"normal", "something", "hogehoge", nil},
		{"nil ID", "", "hogehoge", model.ErrNilID},
	}

	tx := transaction(t)
	for _, c := range candidates {
		sa := model.ServiceAdmin{
			ID:       c.ID,
			Password: c.Password,
		}
		if err := sa.Create(tx); errors.Cause(err) != c.ExpectedErr {
			t.Errorf(`"%s" != "%s"`, err, c.ExpectedErr)
			return
		}
	}
}

func TestServiceAdminLookup(t *testing.T) {
	candidates := []struct {
		ID       string
		Password string
		ExpectedErr error
	}{
		{"", "password", model.ErrNilID},
		{"something", "hogehoge", nil},
	}

	tx := transaction(t)
	for _, c := range candidates {
		sqlmock.ExpectedRows(
			sqlmock.Columns([]string{"id", "password"}),
			sqlmock.ValuesList([]driver.Value{c.ID, c.Password}),
		)

		sa := model.ServiceAdmin{
			ID: c.ID,
		}
		if err := sa.Lookup(tx); errors.Cause(err) != c.ExpectedErr {
			t.Errorf(`"%s" != "%s"`, errors.Cause(err), c.ExpectedErr)
			return
		}
		if c.ExpectedErr != nil {
			continue
		}
		if sa.ID != c.ID {
			t.Errorf(`"%s" != "%s"`, sa.ID, c.ID)
			return
		}
		if sa.Password != c.Password {
			t.Errorf(`"%s" != "%s"`, sa.Password, c.Password)
			return
		}
	}
}

func TestServiceAdminUpdate(t *testing.T) {
	candidates := []struct {
		ID string
		Password string
		ExpectedErr error
	}{
		{"something", "password", nil},
		{"", "password", model.ErrNilID},
		{"something", "", model.ErrNilPasswd},
	}
	tx := transaction(t)
	for _, c:= range candidates {
		sqlmock.ExpectedResult(
			sqlmock.RowsAffected(1),
		)
		sa := model.ServiceAdmin{
			ID: c.ID,
			Password: c.Password,
		}
		if err := sa.Update(tx); errors.Cause(err) != c.ExpectedErr {
			t.Errorf(`"%s" != "%s"`, errors.Cause(err), c.ExpectedErr)
			return
		}
	}
}

func TestServiceAdminDelete(t *testing.T) {
	candidates := []struct {
		ID string
		ExpectedErr error
	}{
		{"something", nil},
		{"", model.ErrNilID},
	}
	tx := transaction(t)
	for _, c := range candidates {
		sqlmock.ExpectedResult(
			sqlmock.RowsAffected(1),
		)

		sa := model.ServiceAdmin{
			ID: c.ID,
		}
		if err := sa.Delete(tx); errors.Cause(err) != c.ExpectedErr {
			t.Errorf(`"%s" != "%s"`, errors.Cause(err), c.ExpectedErr)
			return
		}
	}
}