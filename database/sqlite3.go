package database

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
)

func NewSqliteDb(dbName string) (*sqlx.DB, error) {
	_, dbPath := GetDatabaseFolderPath(dbName)
	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf(`error opening database: %w`, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf(`error connecting to database: %w`, err)
	}

	return db, nil
}

func GetDatabaseFolderPath(dbName string) (migrationsFolder string, databasePath string) {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return "", ""
	}
	// Append a path to the directory
	databasePath = filepath.Join(currentDir, "database", dbName)
	migrationsFolder = filepath.Join(currentDir, "database", "migrations")
	return
}
