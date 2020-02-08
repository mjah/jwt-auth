package auth

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/mjah/jwt-auth/auth/jwt"
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/errors"
	"github.com/mjah/jwt-auth/utils"
)

// SignInDetails ...
type SignInDetails struct {
	Email      string `json:"email" binding:"required" valid:"email"`
	Password   string `json:"password" binding:"required" valid:"length(8|60)"`
	RememberMe bool   `json:"remember_me"`
}

func updateSignInHistory(db *gorm.DB, user *database.User, signInSuccess *bool) *errors.ErrorCode {
	if *signInSuccess {
		if err := db.Model(user).Update(database.User{LastSignin: time.Now().UTC()}).Error; err != nil {
			return errors.New(errors.DatabaseQueryFailed, err)
		}
	} else {
		if err := db.Model(user).Update(database.User{FailedSignin: time.Now().UTC()}).Error; err != nil {
			return errors.New(errors.DatabaseQueryFailed, err)
		}
	}
	return nil
}

// SignIn ...
func (details *SignInDetails) SignIn() (string, string, *errors.ErrorCode) {
	// Validate struct
	if _, err := govalidator.ValidateStruct(details); err != nil {
		return "", "", errors.New(errors.SignInDetailsValidationFailed, err)
	}

	// Get database connection
	db, err := database.GetConnection()
	if err != nil {
		return "", "", errors.New(errors.DatabaseConnectionFailed, err)
	}

	// Declare variables
	condition := &database.User{Email: details.Email}
	user := &database.User{}
	role := &database.Role{}

	// Check email exists
	if err := db.Where(condition).First(user).Error; err != nil {
		if database.IsRecordNotFoundError(err) {
			return "", "", errors.New(errors.EmailDoesNotExist, err)
		}
		return "", "", errors.New(errors.DatabaseQueryFailed, err)
	}

	signInSuccess := false
	defer updateSignInHistory(db, user, &signInSuccess)

	// Check password is correct
	if err := utils.CheckPassword(user.Password, details.Password); err != nil {
		return "", "", errors.New(errors.PasswordCheckFailed, err)
	}

	// Get role name
	if err := db.Where("id = ?", user.RoleID).First(&role).Error; err != nil {
		return "", "", errors.New(errors.DatabaseQueryFailed, err)
	}

	// Issue access token
	accessTokenString, errCode := jwt.IssueAccessToken(user, role.Role)
	if err != nil {
		return "", "", errCode
	}

	// Issue refresh token
	refreshTokenString, errCode := jwt.IssueRefreshToken(user, details.RememberMe)
	if err != nil {
		return "", "", errCode
	}

	signInSuccess = true

	return accessTokenString, refreshTokenString, nil
}
