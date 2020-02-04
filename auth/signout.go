package auth

import (
	"github.com/mjah/jwt-auth/auth/jwt"
	"github.com/mjah/jwt-auth/errors"
)

// SignOut ...
func SignOut(refreshTokenString string, refreshTokenClaims jwt.RefreshTokenClaims) *errors.ErrorCode {
	if errCode := jwt.RevokeRefreshToken(refreshTokenString, refreshTokenClaims); errCode != nil {
		return errCode
	}

	return nil
}

// SignOutAll ...
func SignOutAll(refreshTokenClaims jwt.RefreshTokenClaims) *errors.ErrorCode {
	if errCode := jwt.RevokeRefreshTokenAllBefore(refreshTokenClaims); errCode != nil {
		return errCode
	}

	return nil
}
