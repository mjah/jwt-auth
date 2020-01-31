package auth

import (
	"github.com/asaskevich/govalidator"
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/utils"
)

// SignInDetails ...
type SignInDetails struct {
	Email      string `json:"email" binding:"required" valid:"email"`
	Password   string `json:"password" binding:"required" valid:"length(8|60)"`
	RememberMe bool   `json:"remember_me" binding:"required"`
}

// SignIn ...
func (details *SignInDetails) SignIn() (string, error) {
	if _, err := govalidator.ValidateStruct(details); err != nil {
		return "", err
	}

	db, err := database.GetConnection()
	if err != nil {
		return "", err
	}

	query := &database.User{
		Email: details.Email,
	}

	result := &database.User{}

	// Check email exists
	if err := db.Where(query).First(result).Error; err != nil {
		return "", err
	}

	// Check password is correct
	if err := utils.CheckPassword(result.Password, details.Password); err != nil {
		return "", err
	}

	// Issue token
	tokenString, err := IssueToken()
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
