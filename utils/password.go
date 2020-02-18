package utils

import (
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword ...
func GeneratePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), viper.GetInt("account.password_cost"))
	return string(bytes), err
}

// CheckPassword ...
func CheckPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
