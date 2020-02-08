package auth

import (
	"github.com/mjah/jwt-auth/auth/jwt"
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/errors"
)

// Delete ...
func Delete(refreshTokenClaims jwt.RefreshTokenClaims) *errors.ErrorCode {
	// Get database connection
	db, err := database.GetConnection()
	if err != nil {
		return errors.New(errors.DatabaseConnectionFailed, err)
	}

	// Delete account
	if err := db.Unscoped().Where("id = ?", refreshTokenClaims.UserID).Delete(&database.User{}).Error; err != nil {
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	return nil
}
