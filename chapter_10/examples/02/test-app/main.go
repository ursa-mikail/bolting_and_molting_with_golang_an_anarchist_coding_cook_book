package main

import (
    "encoding/csv"
    "database/sql"
    "fmt"
    "log"
    "os"

    "test-app/libs/db"
    "test-app/utils"
)

func createImage(dbConn *sql.DB, name, contents, team, teamOwner, secretKey, status, statusURL string) {
    sha256Hash := utils.GenerateHash(contents)
    hmacHash := utils.GenerateHMAC(secretKey, contents)
    dataToSign := name + contents + sha256Hash + hmacHash + team + teamOwner + status + statusURL
    statusSignature := utils.GenerateHMAC(secretKey, dataToSign)

    insertSQL := `
    INSERT INTO image (name, contents, sha256, hmac, team, team_owner, status, status_signature, status_url)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

    _, err := dbConn.Exec(insertSQL, name, contents, sha256Hash, hmacHash, team, teamOwner, status, statusSignature, statusURL)
    if err != nil {
        log.Fatal(err)
    }
}

func readImage(dbConn *sql.DB, name string) {
    querySQL := "SELECT * FROM image WHERE name = ?"
    row := dbConn.QueryRow(querySQL, name)

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
            "id":               id,
            "name":             imageName,
            "contents":         contents,
            "sha256":           sha256Hash,
            "hmac":             hmacHash,
            "team":             team,
            "team_owner":       teamOwner,
            "status":           status,
            "status_signature": statusSignature,
            "status_url":       statusURL,
        })
    }
}

func exportToCSV(dbConn *sql.DB, filePath string) {
    querySQL := "SELECT * FROM image"
    rows, err := dbConn.Query(querySQL)
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
    dbConn := db.InitDB()
    defer db.CloseDB(dbConn)

    secretKey := "supersecretkey"

    // Create example
    createImage(dbConn, "example_image", "example_data", "team_a", "owner_a", secretKey, "active", "")

    // Read example
    readImage(dbConn, "example_image")

    // Export to CSV example
    exportToCSV(dbConn, "images_export.csv")
    fmt.Println("Data exported to images_export.csv")
}


/*
% go mod init test-app
% go get github.com/mattn/go-sqlite3
% go run main.go
*/
