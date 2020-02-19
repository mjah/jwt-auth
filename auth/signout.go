package auth

import (
	"github.com/mjah/jwt-auth/auth/jwt"
	"github.com/mjah/jwt-auth/errors"
)

// SignOutDetails holds the details required to sign out the user.
type SignOutDetails struct {
	Claims      jwt.RefreshTokenClaims
	TokenString string
}

// SignOut handles the user sign out.
func (details *SignOutDetails) SignOut() *errors.ErrorCode {
	if errCode := jwt.RevokeRefreshToken(details.Claims.UserID, details.TokenString); errCode != nil {
		return errCode
	}

	return nil
}

// SignOutAll handles the user sign out of all sessions.
func (details *SignOutDetails) SignOutAll() *errors.ErrorCode {
	if errCode := jwt.RevokeRefreshTokenAllBefore(details.Claims.UserID); errCode != nil {
		return errCode
	}

	return nil
}
