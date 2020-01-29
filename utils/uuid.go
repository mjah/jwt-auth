package utils

import "github.com/google/uuid"

// GenerateUUID ...
func GenerateUUID() string {
	return uuid.New().String()
}
