package model_test

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

var mockDB *sql.DB

func TestMain(m *testing.M) {
	db, err := sql.Open("sqlmock", "")
	if err != nil {
		log.Fatalf("%s", err)
	}
	mockDB = db
	code := m.Run()
	mockDB.Close()
	os.Exit(code)
}
