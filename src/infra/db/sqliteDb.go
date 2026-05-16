package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func NewSQLiteConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./database.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
