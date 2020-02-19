package jwt

import (
	"time"

	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/errors"
)

// RevokeRefreshToken revokes the refresh token.
func RevokeRefreshToken(userID uint, tokenString string) *errors.ErrorCode {
	// Get database connection
	db, err := database.GetConnection()
	if err != nil {
		return errors.New(errors.DatabaseConnectionFailed, err)
	}

	// Populate token revocation details
	submitToken := &database.TokenRevocation{
		UserID:       userID,
		RefreshToken: tokenString,
	}

	// Revoke
	if err := db.FirstOrCreate(&database.TokenRevocation{}, submitToken).Error; err != nil {
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	return nil
}

// RevokeRefreshTokenAllBefore revokes all refresh token before now.
func RevokeRefreshTokenAllBefore(userID uint) *errors.ErrorCode {
	// Get database connection
	db, err := database.GetConnection()
	if err != nil {
		return errors.New(errors.DatabaseConnectionFailed, err)
	}

	// Populate token revocation details
	submitToken := &database.TokenRevocation{
		UserID:          userID,
		RevokeAllBefore: time.Now().UTC(),
	}

	// Revoke
	if err := db.Create(submitToken).Error; err != nil {
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	return nil
}

// CheckRefreshTokenRevoked checks if a refresh token has been revoked.
func CheckRefreshTokenRevoked(userID uint, iat int64, tokenString string) *errors.ErrorCode {
	// Get database connection
	db, err := database.GetConnection()
	if err != nil {
		return errors.New(errors.DatabaseConnectionFailed, err)
	}

	// Check if user is active
	user := &database.User{}

	if err := db.Where("id = ?", userID).First(user).Error; err != nil {
		if database.IsRecordNotFoundError(err) {
			return errors.New(errors.UserDoesNotExist, err)
		}
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	if user.IsActive == false {
		return errors.New(errors.UserIsNotActive, nil)
	}

	// Check if token is revoked
	if err := db.Where("user_id = ? AND refresh_token = ?", userID, tokenString).First(&database.TokenRevocation{}).Error; err != nil {
		if !database.IsRecordNotFoundError(err) {
			return errors.New(errors.DatabaseQueryFailed, err)
		}
	} else {
		return errors.New(errors.RefreshTokenIsRevoked, nil)
	}

	if err := db.Where("user_id = ? AND revoke_all_before > ?", userID, time.Unix(iat, 0).UTC()).First(&database.TokenRevocation{}).Error; err != nil {
		if !database.IsRecordNotFoundError(err) {
			return errors.New(errors.DatabaseQueryFailed, err)
		}
	} else {
		return errors.New(errors.RefreshTokenIsRevoked, nil)
	}

	return nil
}
