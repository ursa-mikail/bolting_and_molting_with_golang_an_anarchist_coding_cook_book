package main

import (
    "fmt"

    "project_pack/auth"
    "project_pack/logging"
    "project_pack/db"
    "project_pack/email"
)

func main() {
    // Initialize logger
    logger := logging.NewLogger()

    // Check for user authentication
    userAuthenticated, err := auth.AuthenticateUser("username", "password")
    if err != nil {
        logger.Error("Authentication failed", err)
        return
    }

    if userAuthenticated {
        // Connect to the database
        connection, err := db.Connect("localhost:5432")
        if err != nil {
            logger.Error("Failed to connect to DB", err)
            return
        }
        fmt.Println(connection)
    } else {
        fmt.Println("Authentication failed!")
    }

    // Validate an email
    err = email.ValidateEmail("test@example.com")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Email is valid!")
    }
}
