package model_test

import (
	"database/sql/driver"
	"testing"

	sqlmock "github.com/nasa9084/go-sql-mock"
	"github.com/nasa9084/salamander/salamander/model"
	"github.com/pkg/errors"
)

func TestCorporateCreate(t *testing.T) {
	candidates := []struct {
		Name        string
		ID          string
		ExpectedErr error
	}{
		{"normal", "something", nil},
		{"nil ID", "", model.ErrNilID},
	}

	tx := transaction(t)
	for _, c := range candidates {
		u := model.Corporate{
			ID: c.ID,
		}
		if err := u.Create(tx); errors.Cause(err) != c.ExpectedErr {
			t.Errorf(`"%s" != "%s"`, err, c.ExpectedErr)
			return
		}
	}
}

func TestCorporateLookup(t *testing.T) {
	candidates := []struct {
		ID          string
		ExpectedErr error
	}{
		{"", model.ErrNilID},
		{"something", nil},
	}

	tx := transaction(t)
	for _, c := range candidates {
		sqlmock.ExpectedRows(
			sqlmock.Columns([]string{"id"}),
			sqlmock.ValuesList([]driver.Value{c.ID}),
		)

		u := model.Corporate{
			ID: c.ID,
		}
		if err := u.Lookup(tx); errors.Cause(err) != c.ExpectedErr {
			t.Errorf(`"%s" != "%s"`, errors.Cause(err), c.ExpectedErr)
			return
		}
		if c.ExpectedErr != nil {
			continue
		}
		if u.ID != c.ID {
			t.Errorf(`"%s" != "%s"`, u.ID, c.ID)
			return
		}
	}
}

func TestCorporateUpdate(t *testing.T) {
	candidates := []struct {
		ID          string
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
		u := model.Corporate{
			ID: c.ID,
		}
		if err := u.Update(tx); errors.Cause(err) != c.ExpectedErr {
			t.Errorf(`"%s" != "%s"`, errors.Cause(err), c.ExpectedErr)
			return
		}
	}
}

func TestCorporateDelete(t *testing.T) {
	candidates := []struct {
		ID          string
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

		u := model.Corporate{
			ID: c.ID,
		}
		if err := u.Delete(tx); errors.Cause(err) != c.ExpectedErr {
			t.Errorf(`"%s" != "%s"`, errors.Cause(err), c.ExpectedErr)
			return
		}
	}
}
