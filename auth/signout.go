package auth

import (
	"time"

	"github.com/mjah/jwt-auth/auth/jwt"
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/errors"
)

// SignOut ...
func SignOut(refreshTokenString string, refreshTokenClaims jwt.RefreshTokenClaims) *errors.ErrorCode {
	// Get database connection
	db, err := database.GetConnection()
	if err != nil {
		return errors.New(errors.DatabaseConnectionFailed, nil)
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

// SignOutAll ...
func SignOutAll(refreshTokenClaims jwt.RefreshTokenClaims) *errors.ErrorCode {
	// Get database connection
	db, err := database.GetConnection()
	if err != nil {
		return errors.New(errors.DatabaseConnectionFailed, nil)
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
