package auth

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/errors"
)

// ConfirmDetails ...
type ConfirmDetails struct {
	Email        string `json:"email" binding:"required" valid:"email"`
	ConfirmToken string `json:"confirm_token" binding:"required" valid:"length(36|36)"`
}

// Confirm ...
func (details *ConfirmDetails) Confirm() *errors.ErrorCode {
	// Validate struct
	if _, err := govalidator.ValidateStruct(details); err != nil {
		return errors.New(errors.ConfirmDetailsValidationFailed, err)
	}

	// Get database connection
	db, err := database.GetConnection()
	if err != nil {
		return errors.New(errors.DatabaseConnectionFailed, err)
	}

	// Declare variables
	condition := &database.User{Email: details.Email}
	user := &database.User{}

	// Check email exists
	if err := db.Where(condition).First(user).Error; err != nil {
		if database.IsRecordNotFoundError(err) {
			return errors.New(errors.EmailDoesNotExist, err)
		}
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	// Check if user not confirmed
	if user.IsConfirmed {
		return errors.New(errors.UserAlreadyConfirmed, nil)
	}

	// Check if confirmation token matches
	if user.ConfirmToken != details.ConfirmToken {
		return errors.New(errors.ConfirmTokenDoesNotMatch, nil)
	}

	// Check if confirmation token expired
	if user.ConfirmTokenExpires.Unix() < time.Now().Unix() {
		return errors.New(errors.ConfirmTokenExpired, nil)
	}

	// Confirm user
	if err := db.Model(user).Update(database.User{IsConfirmed: true}).Error; err != nil {
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	return nil
}
