package database

import (
	"database/sql"
	"errors"
	"os"
)

func DatabaseCreation() (*sql.DB, error) {
	databasePass := os.Getenv("DATABASE_PASS")
	databaseUser := os.Getenv("DATABASE_USER")

	db, err := sql.Open("mysql", databaseUser+":"+databasePass+"@tcp(127.0.0.1:3306)/medicare")
	if err != nil {
		return nil, errors.New("unauthorized")
	}
	if db.Ping() != nil {
		return nil, errors.New("unauthorized")
	}

	return db, nil
}
