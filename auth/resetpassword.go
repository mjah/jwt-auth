package auth

import (
	"net/url"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/mjah/jwt-auth/auth/jwt"
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/email"
	"github.com/mjah/jwt-auth/errors"
	"github.com/mjah/jwt-auth/utils"
	"github.com/spf13/viper"
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

	if errCode := jwt.RevokeRefreshTokenAllBefore(user.ID); errCode != nil {
		return errCode
	}

	return nil
}

// SendResetPasswordEmailDetails ...
type SendResetPasswordEmailDetails struct {
	Email string `json:"email" binding:"required" valid:"email"`
}

// SendResetPasswordEmail ...
func (details *SendResetPasswordEmailDetails) SendResetPasswordEmail() *errors.ErrorCode {
	// Validate struct
	if _, err := govalidator.ValidateStruct(details); err != nil {
		return errors.New(errors.SendResetPasswordEmailDetailsValidationFailed, err)
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

	// Update user with reset password token
	user.ResetPassToken = utils.GenerateUUID()
	user.ResetPassTokenExpires = time.Now().Add(viper.GetDuration("account.reset_password_token_expires")).UTC()

	if err := db.Save(user).Error; err != nil {
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	// Send reset password email
	resetPassLink, _ := url.Parse(viper.GetString("account.reset_password_token_endpoint"))
	params := url.Values{}
	params.Add("email", details.Email)
	params.Add("reset_password_token", user.ResetPassToken)
	resetPassLink.RawQuery = params.Encode()

	resetPassEmail := email.ResetPasswordEmailParams{
		ReceipientEmail:   details.Email,
		UserFirstName:     user.FirstName,
		ResetPasswordLink: resetPassLink.String(),
		EmailFromName:     viper.GetString("email.from_name"),
	}

	if err := resetPassEmail.AddToQueue(); err != nil {
		return errors.New(errors.MessageQueueFailed, err)
	}

	return nil
}
