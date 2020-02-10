package auth

import (
	"time"

	"github.com/mjah/jwt-auth/auth/jwt"
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/errors"
	"github.com/spf13/viper"
)

// RefreshTokenDetails ...
type RefreshTokenDetails struct {
	UserID uint
}

// RefreshToken ...
func (details *RefreshTokenDetails) RefreshToken() (string, *errors.ErrorCode) {
	// Get database connection
	db, err := database.GetConnection()
	if err != nil {
		return "", errors.New(errors.DatabaseConnectionFailed, err)
	}

	// Declare variables
	user := &database.User{}
	role := &database.Role{}

	// Get user by ID
	if err := db.Where("id = ?", details.UserID).First(user).Error; err != nil {
		if database.IsRecordNotFoundError(err) {
			return "", errors.New(errors.UserDoesNotExist, err)
		}
		return "", errors.New(errors.DatabaseQueryFailed, err)
	}

	// Get role name
	if err := db.Where("id = ?", user.RoleID).First(&role).Error; err != nil {
		return "", errors.New(errors.DatabaseQueryFailed, err)
	}

	// Issue access token
	atc := jwt.AccessTokenClaims{
		Iat:    time.Now().Unix(),
		Exp:    time.Now().Add(viper.GetDuration("token.access_token_expires")).Unix(),
		UserID: user.ID,
		Role:   role.Role,
	}

	accessTokenString, errCode := atc.IssueAccessToken()
	if errCode != nil {
		return "", errCode
	}

	return accessTokenString, nil
}
