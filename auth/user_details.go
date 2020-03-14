package auth

import (
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/errors"
)

// UserDetailsDetails holds the details required to get the user's details.
type UserDetailsDetails struct {
	UserID uint
}

// ReturnUserDetails holds the details that will be returned.
type ReturnUserDetails struct {
	FirstName        string
	LastName         string
	Email            string
	Username         string
	IsConfirmedEmail bool
}

// UserDetails handles getting the user's details.
func (details *UserDetailsDetails) UserDetails() (*ReturnUserDetails, *errors.ErrorCode) {
	// Get database connection
	db, err := database.GetConnection()
	if err != nil {
		return nil, errors.New(errors.DatabaseConnectionFailed, err.Error())
	}

	// Get user by ID
	user := &database.User{}
	if err := db.Where("id = ?", details.UserID).First(user).Error; err != nil {
		if database.IsRecordNotFoundError(err) {
			return nil, errors.New(errors.UserDoesNotExist, err.Error())
		}
		return nil, errors.New(errors.DatabaseQueryFailed, err.Error())
	}

	// Populate user details to return
	userDetails := &ReturnUserDetails{
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		Email:            user.Email,
		Username:         user.Username,
		IsConfirmedEmail: user.IsConfirmedEmail,
	}

	return userDetails, nil
}
