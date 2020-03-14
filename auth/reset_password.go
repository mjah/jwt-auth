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

// ResetPasswordDetails holds the details required to reset the user's password.
type ResetPasswordDetails struct {
	Email              string `json:"email" binding:"required" valid:"email"`
	ResetPasswordToken string `json:"reset_password_token" binding:"required" valid:"length(36|36)"`
	Password           string `json:"password" binding:"required" valid:"length(8|60)"`
}

// ResetPassword handles the user password reset.
func (details *ResetPasswordDetails) ResetPassword() *errors.ErrorCode {
	// Validate struct
	if _, err := govalidator.ValidateStruct(details); err != nil {
		return errors.New(errors.DetailsInvalid, err.Error())
	}

	// Get database connection
	db, err := database.GetConnection()
	if err != nil {
		return errors.New(errors.DatabaseConnectionFailed, err.Error())
	}

	// Declare variables
	condition := &database.User{Email: details.Email}
	user := &database.User{}

	// Check email exists
	if err := db.Where(condition).First(user).Error; err != nil {
		if database.IsRecordNotFoundError(err) {
			return errors.New(errors.EmailDoesNotExist, err.Error())
		}
		return errors.New(errors.DatabaseQueryFailed, err.Error())
	}

	// Check if reset password token matches
	if user.ResetPassToken != details.ResetPasswordToken {
		return errors.New(errors.UUIDTokenDoesNotMatch, "")
	}

	// Check if reset password token expired
	if user.ResetPassTokenExpires.Unix() < time.Now().Unix() {
		return errors.New(errors.UUIDTokenExpired, "")
	}

	// Update password
	generatedPassword, err := utils.GeneratePassword(details.Password)
	if err != nil {
		return errors.New(errors.PasswordGenerationFailed, err.Error())
	}

	if err := db.Model(user).Update(database.User{Password: generatedPassword}).Error; err != nil {
		return errors.New(errors.DatabaseQueryFailed, err.Error())
	}

	if errCode := jwt.RevokeRefreshTokenAllBefore(user.ID); errCode != nil {
		return errCode
	}

	return nil
}

// SendResetPasswordEmailDetails holds the details required to send the email.
type SendResetPasswordEmailDetails struct {
	Email            string `json:"email" binding:"required" valid:"email"`
	ResetPasswordURL string `json:"reset_password_url" binding:"required" valid:"url"`
}

// SendResetPasswordEmail sends the reset password email to queue.
func (details *SendResetPasswordEmailDetails) SendResetPasswordEmail() *errors.ErrorCode {
	// Validate struct
	if _, err := govalidator.ValidateStruct(details); err != nil {
		return errors.New(errors.DetailsInvalid, err.Error())
	}

	// Get database connection
	db, err := database.GetConnection()
	if err != nil {
		return errors.New(errors.DatabaseConnectionFailed, err.Error())
	}

	// Declare variables
	condition := &database.User{Email: details.Email}
	user := &database.User{}

	// Check email exists
	if err := db.Where(condition).First(user).Error; err != nil {
		if database.IsRecordNotFoundError(err) {
			return errors.New(errors.EmailDoesNotExist, err.Error())
		}
		return errors.New(errors.DatabaseQueryFailed, err.Error())
	}

	// Update user with reset password token
	user.ResetPassToken = utils.GenerateUUID()
	user.ResetPassTokenExpires = time.Now().Add(viper.GetDuration("account.reset_password_token_expires")).UTC()

	if err := db.Save(user).Error; err != nil {
		return errors.New(errors.DatabaseQueryFailed, err.Error())
	}

	// Send reset password email
	resetPassLink, _ := url.Parse(details.ResetPasswordURL)
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
		return errors.New(errors.MessageQueueFailed, err.Error())
	}

	return nil
}
