package main

import (
	"fmt"
	"log"

	"ursa/config"
)

const (
	envFilePath = ".env"
)

func main() {
	cfg, err := config.LoadConfig(envFilePath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Println("App running on port:", cfg.AppPort)
	fmt.Println("Database host:", cfg.DBHost)
	fmt.Println("secret:", cfg.DBPass)
}

/*
2025/06/02 10:41:12 Config loaded success
App running on port: 8080
Database host: localhost
secret: secret_000

go mod init ursa
go mod tidy
go run main.go

""" .env
APP_PORT=8080
DB_HOST=localhost
DB_USER=admin
DB_PASS="secret_000"

# rename this file as `.env` for use by main.go
"""
*/
