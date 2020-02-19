package auth

import (
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/errors"
)

// DeleteDetails holds the details required to delete the user.
type DeleteDetails struct {
	UserID uint
}

// Delete handles the user deletion.
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
