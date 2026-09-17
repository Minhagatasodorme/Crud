package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB() *sql.DB {
	database, err := sql.Open("sqlite3", "crud.db")
	if err != nil {
		log.Fatal(err)
	}

	database.Exec(`
		CREATE TABLE IF NOT EXISTS crud (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nome TEXT NOT NULL,
		email TEXT NOT NULL
		)`,
	)

	return database
}
