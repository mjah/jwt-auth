package auth

import (
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/errors"
)

// DeleteDetails ...
type DeleteDetails struct {
	UserID uint
}

// Delete ...
func (details *DeleteDetails) Delete() *errors.ErrorCode {
	// Get database connection
	db, err := database.GetConnection()
	if err != nil {
		return errors.New(errors.DatabaseConnectionFailed, err)
	}

	// Delete account
	if err := db.Unscoped().Where("id = ?", details.UserID).Delete(&database.User{}).Error; err != nil {
		return errors.New(errors.DatabaseQueryFailed, err)
	}

	return nil
}
