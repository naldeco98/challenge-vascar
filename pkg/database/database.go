package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetDatabaseConnection(pathFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", pathFile)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
