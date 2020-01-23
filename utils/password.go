package utils

import "golang.org/x/crypto/bcrypt"

const (
	// PasswordCost ...
	PasswordCost = 11
)

// GeneratePassword ...
func GeneratePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	return string(bytes), err
}

// CheckPassword ...
func CheckPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
