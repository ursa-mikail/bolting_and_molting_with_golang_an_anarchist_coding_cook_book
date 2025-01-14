package main

import (
    "database/sql"
    "encoding/csv"
    "fmt"
    "crypto/hmac"
    "crypto/sha256"
    "log"
    "os"

    _ "github.com/mattn/go-sqlite3"
)

func initDB() *sql.DB {
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

func generateHash(data string) string {
    hash := sha256.Sum256([]byte(data))
    return fmt.Sprintf("%x", hash)
}

func generateHMAC(secretKey, data string) string {
    h := hmac.New(sha256.New, []byte(secretKey))
    h.Write([]byte(data))
    return fmt.Sprintf("%x", h.Sum(nil))
}

func createImage(db *sql.DB, name, contents, team, teamOwner, secretKey, status, statusURL string) {
    sha256Hash := generateHash(contents)
    hmacHash := generateHMAC(secretKey, contents)
    dataToSign := name + contents + sha256Hash + hmacHash + team + teamOwner + status + statusURL
    statusSignature := generateHMAC(secretKey, dataToSign)

    insertSQL := `
    INSERT INTO image (name, contents, sha256, hmac, team, team_owner, status, status_signature, status_url)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

    _, err := db.Exec(insertSQL, name, contents, sha256Hash, hmacHash, team, teamOwner, status, statusSignature, statusURL)
    if err != nil {
        log.Fatal(err)
    }
}

func readImage(db *sql.DB, name string) {
    querySQL := "SELECT * FROM image WHERE name = ?"
    row := db.QueryRow(querySQL, name)

    var id int
    var imageName, contents, sha256Hash, hmacHash, team, teamOwner, status, statusSignature, statusURL string
    if err := row.Scan(&id, &imageName, &contents, &sha256Hash, &hmacHash, &team, &teamOwner, &status, &statusSignature, &statusURL); err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("No record found.")
        } else {
            log.Fatal(err)
        }
    } else {
        fmt.Printf("Image Record: %+v\n", map[string]interface{}{
            "id":              id,
            "name":            imageName,
            "contents":        contents,
            "sha256":          sha256Hash,
            "hmac":            hmacHash,
            "team":            team,
            "team_owner":      teamOwner,
            "status":          status,
            "status_signature": statusSignature,
            "status_url":      statusURL,
        })
    }
}

func updateImage(db *sql.DB, name, newContents, newTeam, newTeamOwner, secretKey, newStatus, newStatusURL string) {
    sha256Hash := generateHash(newContents)
    hmacHash := generateHMAC(secretKey, newContents)
    dataToSign := name + newContents + sha256Hash + hmacHash + newTeam + newTeamOwner + newStatus + newStatusURL
    statusSignature := generateHMAC(secretKey, dataToSign)

    updateSQL := `
    UPDATE image
    SET contents = ?, sha256 = ?, hmac = ?, team = ?, team_owner = ?, status = ?, status_signature = ?, status_url = ?
    WHERE name = ?`

    _, err := db.Exec(updateSQL, newContents, sha256Hash, hmacHash, newTeam, newTeamOwner, newStatus, statusSignature, newStatusURL, name)
    if err != nil {
        log.Fatal(err)
    }
}

func deleteImage(db *sql.DB, name string) {
    deleteSQL := "DELETE FROM image WHERE name = ?"
    _, err := db.Exec(deleteSQL, name)
    if err != nil {
        log.Fatal(err)
    }
}

func exportToCSV(db *sql.DB, filePath string) {
    querySQL := "SELECT * FROM image"
    rows, err := db.Query(querySQL)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    file, err := os.Create(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    columns, err := rows.Columns()
    if err != nil {
        log.Fatal(err)
    }
    writer.Write(columns)

    record := make([]interface{}, len(columns))
    recordPointers := make([]interface{}, len(columns))
    for i := range record {
        recordPointers[i] = &record[i]
    }

    for rows.Next() {
        if err := rows.Scan(recordPointers...); err != nil {
            log.Fatal(err)
        }

        rowData := make([]string, len(columns))
        for i, value := range record {
            rowData[i] = fmt.Sprintf("%v", value)
        }
        writer.Write(rowData)
    }
}

func main() {
    db := initDB()
    defer db.Close()

    secretKey := "supersecretkey"

    // Create example
    createImage(db, "example_image", "example_data", "team_a", "owner_a", secretKey, "active", "")

    // Read example
    readImage(db, "example_image")

    // Update example
    updateImage(db, "example_image", "new_data", "team_b", "owner_b", secretKey, "suspended", "https://example.com/cve-details")

    // Export to CSV example
    exportToCSV(db, "images_export.csv")
    fmt.Println("Data exported to images_export.csv")

    // Delete example
    deleteImage(db, "example_image")
}


/*
% go mod init test-app
% go get github.com/mattn/go-sqlite3
% go run main.go
*/
