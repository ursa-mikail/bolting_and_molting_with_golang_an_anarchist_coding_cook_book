package db

import "fmt"

// Connect simulates a connection to the database
func Connect(connectionString string) (string, error) {
	if connectionString == "localhost:5432" {
		return "Connection successful", nil
	}
	return "", fmt.Errorf("unable to connect to database")
}

