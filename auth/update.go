package auth

import (
	"github.com/asaskevich/govalidator"
	"github.com/mjah/jwt-auth/auth/jwt"
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/errors"
	"github.com/mjah/jwt-auth/utils"
)

// UpdateDetails ...
type UpdateDetails struct {
	Claims    jwt.RefreshTokenClaims
	Email     string `json:"email" valid:"optional,email"`
	Username  string `json:"username" valid:"optional,length(3|40)"`
	Password  string `json:"password" valid:"optional,length(8|60)"`
	FirstName string `json:"first_name" valid:"optional,length(1|32)"`
	LastName  string `json:"last_name" valid:"optional,length(1|32)"`
}

// Update ...
func (details *UpdateDetails) Update() *errors.ErrorCode {
	// Validate struct
	if _, err := govalidator.ValidateStruct(details); err != nil {
		return errors.New(errors.UpdateDetailsValidationFailed, err)
	}

	// Get database connection
	db, err := database.GetConnection()
	if err != nil {
		return errors.New(errors.DatabaseConnectionFailed, err)
	}

	// Get user by ID
	user := &database.User{}
	if err := db.Where("id = ?", details.Claims.UserID).First(user).Error; err != nil {
		if database.IsRecordNotFoundError(err) {
			return errors.New(errors.UserDoesNotExist, err)
		}
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	// Email
	emailAlreadyExists := false
	if details.Email != "" {
		// Check email already exists
		emailAlreadyExists = true
		if err := db.Where(&database.User{Email: details.Email}).First(&database.User{}).Error; err != nil {
			if database.IsRecordNotFoundError(err) {
				emailAlreadyExists = false
				user.Email = details.Email
			} else {
				return errors.New(errors.DatabaseQueryFailed, err)
			}
		}
	}

	// Username
	usernameAlreadyExists := false
	if details.Username != "" {
		// Check username already exists
		usernameAlreadyExists = true
		if err := db.Where(&database.User{Username: details.Username}).First(&database.User{}).Error; err != nil {
			if database.IsRecordNotFoundError(err) {
				usernameAlreadyExists = false
				user.Username = details.Username
			} else {
				return errors.New(errors.DatabaseQueryFailed, err)
			}
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

	// Password
	if details.Password != "" {
		// Generate password
		generatedPassword, err := utils.GeneratePassword(details.Password)
		if err != nil {
			return errors.New(errors.PasswordGenerationFailed, err)
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
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	// Revoke refresh token all before on password change
	if details.Password != "" {
		if errCode := jwt.RevokeRefreshTokenAllBefore(details.Claims); errCode != nil {
			return errCode
		}
	}

	return nil
}
