package storage

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func GetDatabaseConnection(pathFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", pathFile)
	if err != nil {
		return nil, err
	}
	// Check if database exist in pathFile
	if _, err := os.Stat(pathFile); err != nil {
		return nil, err
	}
	return db, nil
}
