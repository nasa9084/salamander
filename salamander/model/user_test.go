package model_test

import (
	"database/sql/driver"
	"testing"

	sqlmock "github.com/nasa9084/go-sql-mock"
	"github.com/nasa9084/salamander/salamander/model"
	"github.com/pkg/errors"
)

func TestUserCreate(t *testing.T) {
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
		u := model.User{
			ID:       c.ID,
			Password: c.Password,
		}
		if err := u.Create(tx); errors.Cause(err) != c.ExpectedErr {
			t.Errorf(`"%s" != "%s"`, err, c.ExpectedErr)
			return
		}
	}
}

func TestUserLookup(t *testing.T) {
	candidates := []struct {
		ID          string
		Password    string
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

		u := model.User{
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
		if u.Password != c.Password {
			t.Errorf(`"%s" != "%s"`, u.Password, c.Password)
			return
		}
	}
}

func TestUserUpdate(t *testing.T) {
	candidates := []struct {
		ID          string
		Password    string
		ExpectedErr error
	}{
		{"something", "password", nil},
		{"", "password", model.ErrNilID},
		{"something", "", model.ErrNilPasswd},
	}
	tx := transaction(t)
	for _, c := range candidates {
		sqlmock.ExpectedResult(
			sqlmock.RowsAffected(1),
		)
		u := model.User{
			ID:       c.ID,
			Password: c.Password,
		}
		if err := u.Update(tx); errors.Cause(err) != c.ExpectedErr {
			t.Errorf(`"%s" != "%s"`, errors.Cause(err), c.ExpectedErr)
			return
		}
	}
}

func TestUserDelete(t *testing.T) {
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

		u := model.User{
			ID: c.ID,
		}
		if err := u.Delete(tx); errors.Cause(err) != c.ExpectedErr {
			t.Errorf(`"%s" != "%s"`, errors.Cause(err), c.ExpectedErr)
			return
		}
	}
}
