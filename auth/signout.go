package auth

import (
	"github.com/mjah/jwt-auth/auth/jwt"
	"github.com/mjah/jwt-auth/errors"
)

// SignOutDetails ...
type SignOutDetails struct {
	Claims      jwt.RefreshTokenClaims
	TokenString string
}

// SignOut ...
func (details *SignOutDetails) SignOut() *errors.ErrorCode {
	if errCode := jwt.RevokeRefreshToken(details.TokenString, details.Claims); errCode != nil {
		return errCode
	}

	return nil
}

// SignOutAll ...
func (details *SignOutDetails) SignOutAll() *errors.ErrorCode {
	if errCode := jwt.RevokeRefreshTokenAllBefore(details.Claims); errCode != nil {
		return errCode
	}

	return nil
}
