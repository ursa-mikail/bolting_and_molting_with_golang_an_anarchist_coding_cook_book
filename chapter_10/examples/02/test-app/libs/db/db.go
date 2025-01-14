package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the database and creates the required table if not exists.
func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "images.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS image (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		contents TEXT NOT NULL,
		sha256 TEXT NOT NULL,
		hmac TEXT NOT NULL,
		team TEXT NOT NULL,
		team_owner TEXT NOT NULL,
		status TEXT CHECK(status IN ('active', 'suspended', 'revoked')) NOT NULL DEFAULT 'active',
		status_signature TEXT,
		status_url TEXT
	)`

	if _, err := db.Exec(createTable); err != nil {
		log.Fatal(err)
	}

	return db
}

// CloseDB safely closes the database connection.
func CloseDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
