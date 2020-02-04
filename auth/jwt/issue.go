package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mjah/jwt-auth/database"
	"github.com/spf13/viper"
)

const issuer = "jwt-auth"

// IssueAccessToken ...
func IssueAccessToken(user *database.User, role string) (string, error) {
	expireIn := viper.GetDuration("token.access_token_expires")
	accessTokenClaims := jwt.MapClaims{
		"sub":          "access",
		"iss":          issuer,
		"iat":          time.Now().Unix(),
		"exp":          time.Now().Add(expireIn).Unix(),
		"user_id":      user.ID,
		"role_name":    role,
		"is_active":    user.IsActive,
		"is_confirmed": user.IsConfirmed,
	}
	return jwt.NewWithClaims(jwt.SigningMethodRS256, accessTokenClaims).SignedString(privateKey)
}

// IssueRefreshToken ...
func IssueRefreshToken(user *database.User, extendedRefresh bool) (string, error) {
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
	return jwt.NewWithClaims(jwt.SigningMethodRS256, refreshTokenClaims).SignedString(privateKey)
}
