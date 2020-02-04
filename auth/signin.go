package auth

import (
	"github.com/asaskevich/govalidator"
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

	// to-do: update last_signin or failed_signin field
	//        check if locked
	//        check if active

	return accessTokenString, refreshTokenString, nil
}
