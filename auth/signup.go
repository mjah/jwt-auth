package auth

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/errors"
	"github.com/mjah/jwt-auth/utils"
	"github.com/spf13/viper"
)

// SignUpDetails ...
type SignUpDetails struct {
	Email     string `json:"email" binding:"required" valid:"email"`
	Username  string `json:"username" binding:"required" valid:"length(3|40)"`
	Password  string `json:"password" binding:"required" valid:"length(8|60)"`
	FirstName string `json:"first_name" binding:"required" valid:"length(1|32)"`
	LastName  string `json:"last_name" binding:"required" valid:"length(1|32)"`
}

// SignUp ...
func (details *SignUpDetails) SignUp() *errors.ErrorCode {
	// Validate details
	if _, err := govalidator.ValidateStruct(details); err != nil {
		return errors.New(errors.SignUpDetailsValidationFailed, err)
	}

	// Get database connection
	db, err := database.GetConnection()
	if err != nil {
		return errors.New(errors.DatabaseConnectionFailed, err)
	}

	// Check email already exists
	emailAlreadyExists := true
	if err := db.Where(&database.User{Email: details.Email}).First(&database.User{}).Error; err != nil {
		if database.IsRecordNotFoundError(err) {
			emailAlreadyExists = false
		} else {
			return errors.New(errors.DatabaseQueryFailed, err)
		}
	}

	// Check username already exists
	usernameAlreadyExists := true
	if err := db.Where(&database.User{Username: details.Username}).First(&database.User{}).Error; err != nil {
		if database.IsRecordNotFoundError(err) {
			usernameAlreadyExists = false
		} else {
			return errors.New(errors.DatabaseQueryFailed, err)
		}
	}

	// If email and/or username exists, return error
	if emailAlreadyExists && usernameAlreadyExists {
		return errors.New(errors.EmailAndUsernameAlreadyExists, nil)
	} else if emailAlreadyExists {
		return errors.New(errors.EmailAlreadyExists, nil)
	} else if usernameAlreadyExists {
		return errors.New(errors.UsernameAlreadyExists, nil)
	}

	// Get default role ID
	role := &database.Role{}
	if err := db.Where("role = ?", viper.GetString("roles.default")).First(&role).Error; err != nil {
		if database.IsRecordNotFoundError(err) {
			return errors.New(errors.DefaultRoleDoesNotExist, err)
		}
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	// Generate password
	generatedPassword, err := utils.GeneratePassword(details.Password)
	if err != nil {
		return errors.New(errors.PasswordGenerationFailed, err)
	}

	// Populate user details to be submitted
	submitUser := &database.User{
		RoleID:              role.ID,
		Email:               details.Email,
		Username:            details.Username,
		Password:            generatedPassword,
		FirstName:           details.FirstName,
		LastName:            details.LastName,
		ConfirmToken:        utils.GenerateUUID(),
		ConfirmTokenExpires: time.Now().Add(time.Hour * 24).UTC(),
	}

	// Execute query
	if err := db.Create(submitUser).Error; err != nil {
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	return nil
}
