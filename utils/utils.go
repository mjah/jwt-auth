// Package utils is a collection of helpful utilities for common actions.
package utils

import (
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// GenerateUUID returns a random UUID string.
func GenerateUUID() string {
	return uuid.New().String()
}

// GeneratePassword returns the bcrypt hash string of the given password at the
// configured cost.
func GeneratePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), viper.GetInt("account.password_cost"))
	return string(bytes), err
}

// CheckPassword compares a bcrypt hashed password with its possible plaintext
// equivalent. Returns nil on success, or an error on failure.
func CheckPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
