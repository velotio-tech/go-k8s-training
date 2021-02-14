package cryptography

import (
	"fmt"

	// for hashing password
	"golang.org/x/crypto/bcrypt"
)

// Hash - encrypts the password and returns it
func Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", hashedPassword), nil
}

// ValidatePassword - validates password with existing hash in DB
func ValidatePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
