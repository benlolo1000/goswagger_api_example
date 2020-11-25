package userDatabase

import (
	"database/sql"
	"errors"
	"os"
)

func DatabaseCreation() (*sql.DB, error) {
	databasePass := os.Getenv("DATABASE_PASS")
	databaseUser := os.Getenv("DATABASE_USER")

	userDb, err := sql.Open("mysql", databaseUser+":"+databasePass+"@tcp(127.0.0.1:3306)/medicareUsers")

	if err != nil {
		return nil, errors.New("unauthorized")
	}
	if userDb.Ping() != nil {
		return nil, errors.New("unauthorized")
	}

	return userDb, nil
}
