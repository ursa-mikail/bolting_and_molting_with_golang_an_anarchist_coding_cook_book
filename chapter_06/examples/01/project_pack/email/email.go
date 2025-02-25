package email

import (
	"github.com/go-playground/validator/v10"
	"fmt"
)

// ValidateEmail checks if the email address is valid
func ValidateEmail(email string) error {
	validate := validator.New()
	err := validate.Var(email, "required,email")
	if err != nil {
		return fmt.Errorf("invalid email address")
	}
	return nil
}

