package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func SetupSQLite(filePath string) *sql.DB {
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		log.Fatal(err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL
	);`
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

	return db
}
