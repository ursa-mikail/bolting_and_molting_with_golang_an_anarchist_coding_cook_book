package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the SQLite database and creates the table if it doesn't exist.
func InitDB(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
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
