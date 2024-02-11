package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"
)

func NewSqliteDb(dbName string) (*sqlx.DB, error) {
	_, dbPath := GetDatabaseFolderPath(dbName)

	// Register SQLite3 driver explicitly
	sql.Register("sqlite3_driver", &sqlite3.SQLiteDriver{})

	fmt.Println(dbPath)

	db, err := sqlx.Open("sqlite3_driver", dbPath)
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
