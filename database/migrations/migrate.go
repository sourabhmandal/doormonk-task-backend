package main

import (
	"database/sql"
	"dmbackend/database"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", "go run database/migrations/migrate.go", " [up|down]")
		return
	}

	direction := os.Args[1]
	if direction != "up" && direction != "down" {
		fmt.Println("Invalid direction. Usage: ", "go run database/migrations/migrate.go", " [up|down]")
		return
	}

	migrationsFolder, databasePath := database.GetDatabaseFolderPath("doormonk.db")

	fmt.Println(migrationsFolder)
	fmt.Println(databasePath)

	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsFolder,
		"sqlite3", driver)
	if err != nil {
		panic(err)
	}
	if direction == "up" {
		log.Println("migrating up")

		if err := m.Up(); err != nil {
			panic(err)
		}
	}
	if direction == "down" {
		log.Println("migrating down")
		if err := m.Down(); err != nil {
			panic(err)
		}
	}

}
