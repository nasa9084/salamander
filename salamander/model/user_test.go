package model_test

import (
	"database/sql/driver"
	"testing"

	sqlmock "github.com/nasa9084/go-sql-mock"
	"github.com/nasa9084/salamander/salamander/model"
	"github.com/pkg/errors"
)

/* TODO: refactoring */

func userTestCase(drop string) *model.User {
	u := &model.User{
		ID:          "hoge",
		Password:    "hoge",
		DisplayName: "hoge",
		Email:       "hoge",
	}
	switch drop {
	case "id":
		u.ID = ""
	case "password":
		u.Password = ""
	case "displayname":
		u.DisplayName = ""
	case "email":
		u.Email = ""
	}
	return u
}

func TestUserCreate(t *testing.T) {
	candidates := []struct {
		Name        string
		User        *model.User
		ExpectedErr error
	}{
		{"normal", userTestCase(""), nil},
		{"nil ID", userTestCase("id"), model.ErrNilID},
		{"nil Password", userTestCase("password"), model.ErrNilPasswd},
	}

	tx := transaction(t)
	for _, c := range candidates {
		u := model.User{
			ID:       c.User.ID,
			Password: c.User.Password,
		}
		if err := u.Create(tx); errors.Cause(err) != c.ExpectedErr {
			t.Errorf(`"%s" != "%s"`, err, c.ExpectedErr)
			return
		}
	}
}

func TestUserLookup(t *testing.T) {
	candidates := []struct {
		User        *model.User
		ExpectedErr error
	}{
		{userTestCase("id"), model.ErrNilID},
		{userTestCase(""), nil},
	}

	tx := transaction(t)
	for _, c := range candidates {
		sqlmock.ExpectedRows(
			sqlmock.Columns([]string{"id", "password", "display_name", "email"}),
			sqlmock.ValuesList([]driver.Value{
				c.User.ID,
				c.User.Password,
				c.User.DisplayName,
				c.User.Email,
			}),
		)

		u := model.User{
			ID: c.User.ID,
		}
		if err := u.Lookup(tx); errors.Cause(err) != c.ExpectedErr {
			t.Errorf(`"%s" != "%s"`, errors.Cause(err), c.ExpectedErr)
			return
		}
		if c.ExpectedErr != nil {
			continue
		}
		if u.ID != c.User.ID {
			t.Errorf(`"%s" != "%s"`, u.ID, c.User.ID)
			return
		}
		if u.Password != c.User.Password {
			t.Errorf(`"%s" != "%s"`, u.Password, c.User.Password)
			return
		}
	}
}

func TestUserUpdate(t *testing.T) {
	candidates := []struct {
		User        *model.User
		ExpectedErr error
	}{
		{userTestCase(""), nil},
		{userTestCase("id"), model.ErrNilID},
		{userTestCase("password"), model.ErrNilPasswd},
	}
	tx := transaction(t)
	for _, c := range candidates {
		sqlmock.ExpectedResult(
			sqlmock.RowsAffected(1),
		)
		u := model.User{
			ID:       c.User.ID,
			Password: c.User.Password,
		}
		if err := u.Update(tx); errors.Cause(err) != c.ExpectedErr {
			t.Errorf(`"%s" != "%s"`, errors.Cause(err), c.ExpectedErr)
			return
		}
	}
}

func TestUserDelete(t *testing.T) {
	candidates := []struct {
		User        *model.User
		ExpectedErr error
	}{
		{userTestCase(""), nil},
		{userTestCase("id"), model.ErrNilID},
	}
	tx := transaction(t)
	for _, c := range candidates {
		sqlmock.ExpectedResult(
			sqlmock.RowsAffected(1),
		)

		u := model.User{
			ID: c.User.ID,
		}
		if err := u.Delete(tx); errors.Cause(err) != c.ExpectedErr {
			t.Errorf(`"%s" != "%s"`, errors.Cause(err), c.ExpectedErr)
			return
		}
	}
}
