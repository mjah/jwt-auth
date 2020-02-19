package auth

import (
	"net/url"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/email"
	"github.com/mjah/jwt-auth/errors"
	"github.com/mjah/jwt-auth/utils"
	"github.com/spf13/viper"
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
		return errors.New(errors.DetailsInvalid, err)
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
		return errors.New(errors.UUIDTokenDoesNotMatch, nil)
	}

	// Check if confirmation token expired
	if user.ConfirmTokenExpires.Unix() < time.Now().Unix() {
		return errors.New(errors.UUIDTokenExpired, nil)
	}

	// Confirm user
	if err := db.Model(user).Update(database.User{IsConfirmed: true}).Error; err != nil {
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	return nil
}

// SendConfirmEmailDetails ...
type SendConfirmEmailDetails struct {
	Email string `json:"email" binding:"required" valid:"email"`
}

// SendConfirmEmail ...
func (details *SendConfirmEmailDetails) SendConfirmEmail() *errors.ErrorCode {
	// Validate details
	if _, err := govalidator.ValidateStruct(details); err != nil {
		return errors.New(errors.DetailsInvalid, err)
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

	// Update user with confirm token
	user.ConfirmToken = utils.GenerateUUID()
	user.ConfirmTokenExpires = time.Now().Add(viper.GetDuration("account.confirm_token_expires")).UTC()

	if err := db.Save(user).Error; err != nil {
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	// Send confirm email
	confirmLink, _ := url.Parse(viper.GetString("account.confirm_token_endpoint"))
	params := url.Values{}
	params.Add("email", details.Email)
	params.Add("confirm_token", user.ConfirmToken)
	confirmLink.RawQuery = params.Encode()

	confirmEmail := email.ConfirmEmailParams{
		ReceipientEmail:  details.Email,
		UserFirstName:    user.FirstName,
		ConfirmationLink: confirmLink.String(),
		EmailFromName:    viper.GetString("email.from_name"),
	}

	if err := confirmEmail.AddToQueue(); err != nil {
		return errors.New(errors.MessageQueueFailed, err)
	}

	return nil
}
