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

// ConfirmEmailDetails holds the details required to confirm the user's email.
type ConfirmEmailDetails struct {
	Email             string `json:"email" binding:"required" valid:"email"`
	ConfirmEmailToken string `json:"confirm_email_token" binding:"required" valid:"length(36|36)"`
}

// ConfirmEmail handles the email confirmation.
func (details *ConfirmEmailDetails) ConfirmEmail() *errors.ErrorCode {
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

	// Check if email not confirmed
	if user.IsConfirmedEmail {
		return errors.New(errors.EmailAlreadyConfirmed, nil)
	}

	// Check if confirmation token matches
	if user.ConfirmEmailToken != details.ConfirmEmailToken {
		return errors.New(errors.UUIDTokenDoesNotMatch, nil)
	}

	// Check if confirmation token expired
	if user.ConfirmEmailTokenExpires.Unix() < time.Now().Unix() {
		return errors.New(errors.UUIDTokenExpired, nil)
	}

	// Confirm email
	if err := db.Model(user).Update(database.User{IsConfirmedEmail: true}).Error; err != nil {
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	return nil
}

// SendConfirmEmailDetails holds the details required to send the email.
type SendConfirmEmailDetails struct {
	Email           string `json:"email" binding:"required" valid:"email"`
	ConfirmEmailURL string `json:"confirm_email_url" binding:"required" valid:"url"`
}

// SendConfirmEmail sends confirm email email to queue.
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

	// Check if email not confirmed
	if user.IsConfirmedEmail {
		return errors.New(errors.EmailAlreadyConfirmed, nil)
	}

	// Update user with confirm email token
	user.ConfirmEmailToken = utils.GenerateUUID()
	user.ConfirmEmailTokenExpires = time.Now().Add(viper.GetDuration("account.confirm_token_expires")).UTC()

	if err := db.Save(user).Error; err != nil {
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	// Send confirm email email
	confirmLink, _ := url.Parse(details.ConfirmEmailURL)
	params := url.Values{}
	params.Add("email", details.Email)
	params.Add("confirm_token", user.ConfirmEmailToken)
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
