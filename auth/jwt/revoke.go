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
		RevokeAllBefore: time.Now(),
	}

	// Revoke
	if err := db.Create(submitToken).Error; err != nil {
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	return nil
}
