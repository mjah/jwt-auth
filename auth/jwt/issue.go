package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/errors"
	"github.com/spf13/viper"
)

const issuer = "jwt-auth"

// IssueAccessToken ...
func IssueAccessToken(user *database.User, role string) (string, *errors.ErrorCode) {
	expireIn := viper.GetDuration("token.access_token_expires")

	accessTokenClaims := jwt.MapClaims{
		"sub":     "access",
		"iss":     issuer,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(expireIn).Unix(),
		"user_id": user.ID,
		"role":    role,
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodRS256, accessTokenClaims).SignedString(privateKey)
	if err != nil {
		return "", errors.New(errors.AccessTokenIssueFailed, err)
	}

	return tokenString, nil
}

// IssueRefreshToken ...
func IssueRefreshToken(user *database.User, extendedRefresh bool) (string, *errors.ErrorCode) {
	expireIn := viper.GetDuration("token.refresh_token_expires")

	if extendedRefresh {
		expireIn = viper.GetDuration("token.refresh_token_expires_extended")
	}

	refreshTokenClaims := jwt.MapClaims{
		"sub":     "refresh",
		"iss":     issuer,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(expireIn).Unix(),
		"user_id": user.ID,
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshTokenClaims).SignedString(privateKey)
	if err != nil {
		return "", errors.New(errors.RefreshTokenIssueFailed, err)
	}

	return tokenString, nil
}
