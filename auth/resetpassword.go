package auth

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/errors"
	"github.com/mjah/jwt-auth/utils"
)

// ResetPasswordDetails ...
type ResetPasswordDetails struct {
	Email              string `json:"email" binding:"required" valid:"email"`
	ResetPasswordToken string `json:"reset_password_token" binding:"required" valid:"length(36|36)"`
	Password           string `json:"password" binding:"required" valid:"length(8|60)"`
}

// ResetPassword ...
func (details *ResetPasswordDetails) ResetPassword() *errors.ErrorCode {
	// Validate struct
	if _, err := govalidator.ValidateStruct(details); err != nil {
		return errors.New(errors.ResetPasswordDetailsValidationFailed, err)
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

	// Check if reset password token matches
	if user.ResetPassToken != details.ResetPasswordToken {
		return errors.New(errors.ResetPasswordTokenDoesNotMatch, nil)
	}

	// Check if reset password token expired
	if user.ResetPassTokenExpires.Unix() < time.Now().Unix() {
		return errors.New(errors.ResetPasswordTokenExpired, nil)
	}

	// Update password
	generatedPassword, err := utils.GeneratePassword(details.Password)
	if err != nil {
		return errors.New(errors.PasswordGenerationFailed, err)
	}

	if err := db.Model(user).Update(database.User{Password: generatedPassword}).Error; err != nil {
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	// to-do: revoke all refresh token before

	return nil
}
