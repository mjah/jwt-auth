package jwt

import (
	"time"

	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/errors"
)

// RevokeRefreshToken ...
func RevokeRefreshToken(refreshTokenString string, refreshTokenClaims RefreshTokenClaims) *errors.ErrorCode {
	// Get database connection
	db, err := database.GetConnection()
	if err != nil {
		return errors.New(errors.DatabaseConnectionFailed, err)
	}

	// Populate token revocation details
	submitToken := &database.TokenRevocation{
		UserID:       refreshTokenClaims.UserID,
		RefreshToken: refreshTokenString,
	}

	// Revoke
	if err := db.FirstOrCreate(&database.TokenRevocation{}, submitToken).Error; err != nil {
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	return nil
}

// RevokeRefreshTokenAllBefore ...
func RevokeRefreshTokenAllBefore(refreshTokenClaims RefreshTokenClaims) *errors.ErrorCode {
	// Get database connection
	db, err := database.GetConnection()
	if err != nil {
		return errors.New(errors.DatabaseConnectionFailed, err)
	}

	// Populate token revocation details
	submitToken := &database.TokenRevocation{
		UserID:          refreshTokenClaims.UserID,
		RevokeAllBefore: time.Now().UTC(),
	}

	// Revoke
	if err := db.Create(submitToken).Error; err != nil {
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	return nil
}

// CheckRefreshTokenRevoked ...
func CheckRefreshTokenRevoked(claims RefreshTokenClaims, tokenString string) *errors.ErrorCode {
	// Get database connection
	db, err := database.GetConnection()
	if err != nil {
		return errors.New(errors.DatabaseConnectionFailed, err)
	}

	// Check if user is active
	user := &database.User{}

	if err := db.Where("id = ?", claims.UserID).First(user).Error; err != nil {
		if database.IsRecordNotFoundError(err) {
			return errors.New(errors.UserDoesNotExist, err)
		}
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	if user.IsActive == false {
		return errors.New(errors.UserIsNotActive, nil)
	}

	// Check if token is revoked
	if err := db.Where("user_id = ? AND refresh_token = ?", claims.UserID, tokenString).First(&database.TokenRevocation{}).Error; err != nil {
		if !database.IsRecordNotFoundError(err) {
			return errors.New(errors.DatabaseQueryFailed, err)
		}
	} else {
		return errors.New(errors.RefreshTokenIsRevoked, nil)
	}

	if err := db.Where("user_id = ? AND revoke_all_before > ?", claims.UserID, time.Unix(claims.Iat, 0).UTC()).First(&database.TokenRevocation{}).Error; err != nil {
		if !database.IsRecordNotFoundError(err) {
			return errors.New(errors.DatabaseQueryFailed, err)
		}
	} else {
		return errors.New(errors.RefreshTokenIsRevoked, nil)
	}

	return nil
}
