package auth

import "fmt"

// AuthenticateUser is responsible for checking the user's credentials
func AuthenticateUser(username, password string) (bool, error) {
	if username == "admin" && password == "secret" {
		return true, nil
	}
	return false, fmt.Errorf("invalid credentials")
}

