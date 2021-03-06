package auth

import (
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/mjah/jwt-auth/auth/jwt"
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/errors"
	"github.com/mjah/jwt-auth/utils"
)

// UpdateDetails holds the details required to update the user.
type UpdateDetails struct {
	Claims    jwt.RefreshTokenClaims
	Email     string `json:"email" valid:"optional,email"`
	Username  string `json:"username" valid:"optional,length(3|40)"`
	Password  string `json:"password" valid:"optional,length(8|60)"`
	FirstName string `json:"first_name" valid:"optional,length(1|32)"`
	LastName  string `json:"last_name" valid:"optional,length(1|32)"`
}

// Update handles the user update.
func (details *UpdateDetails) Update() *errors.ErrorCode {
	// Validate struct
	if _, err := govalidator.ValidateStruct(details); err != nil {
		return errors.New(errors.DetailsInvalid, err.Error())
	}

	// Get database connection
	db, err := database.GetConnection()
	if err != nil {
		return errors.New(errors.DatabaseConnectionFailed, err.Error())
	}

	// Get user by ID
	user := &database.User{}
	if err := db.Where("id = ?", details.Claims.UserID).First(user).Error; err != nil {
		if database.IsRecordNotFoundError(err) {
			return errors.New(errors.UserDoesNotExist, err.Error())
		}
		return errors.New(errors.DatabaseQueryFailed, err.Error())
	}

	// Email
	emailAlreadyExists := false
	details.Email = strings.ToLower(details.Email)
	if details.Email != "" && details.Email != user.Email {
		// Check email already exists
		emailAlreadyExists = true
		if err := db.Where(&database.User{Email: details.Email}).First(&database.User{}).Error; err != nil {
			if database.IsRecordNotFoundError(err) {
				emailAlreadyExists = false
				user.Email = details.Email
				user.IsConfirmedEmail = false
			} else {
				return errors.New(errors.DatabaseQueryFailed, err.Error())
			}
		}
	}

	// Username
	usernameAlreadyExists := false
	if details.Username != "" && details.Username != user.Username {
		// Check username already exists
		usernameAlreadyExists = true
		if err := db.Where(&database.User{Username: details.Username}).First(&database.User{}).Error; err != nil {
			if database.IsRecordNotFoundError(err) {
				usernameAlreadyExists = false
				user.Username = details.Username
			} else {
				return errors.New(errors.DatabaseQueryFailed, err.Error())
			}
		}
	}

	// If email and/or username exists, return error
	if emailAlreadyExists && usernameAlreadyExists {
		return errors.New(errors.EmailAndUsernameAlreadyExists, "")
	} else if emailAlreadyExists {
		return errors.New(errors.EmailAlreadyExists, "")
	} else if usernameAlreadyExists {
		return errors.New(errors.UsernameAlreadyExists, "")
	}

	// Password
	if details.Password != "" {
		// Generate password
		generatedPassword, err := utils.GeneratePassword(details.Password)
		if err != nil {
			return errors.New(errors.PasswordGenerationFailed, err.Error())
		}
		user.Password = generatedPassword
	}

	// First name
	if details.FirstName != "" {
		user.FirstName = details.FirstName
	}

	// Last name
	if details.LastName != "" {
		user.LastName = details.LastName
	}

	// Update user
	if err := db.Save(user).Error; err != nil {
		return errors.New(errors.DatabaseQueryFailed, err.Error())
	}

	// Revoke refresh token all before on password change
	if details.Password != "" {
		if errCode := jwt.RevokeRefreshTokenAllBefore(details.Claims.UserID); errCode != nil {
			return errCode
		}
	}

	return nil
}
