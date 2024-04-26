package common

import (
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

// Check if a given email is valid (using Go's standard library)
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// Hash a given password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
