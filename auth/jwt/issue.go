package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

const issuer = "jwt-auth"

// IssueAccessToken ...
func IssueAccessToken(userID uint, role string) (string, error) {
	expireIn := viper.GetDuration("token.access_token_expires")
	accessTokenClaims := jwt.MapClaims{
		"iss":     issuer,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(expireIn).Unix(),
		"user_id": userID,
		"role":    role,
	}
	return jwt.NewWithClaims(jwt.SigningMethodRS256, accessTokenClaims).SignedString(privateKey)
}

// IssueRefreshToken ...
func IssueRefreshToken(userID uint, extendedRefresh bool) (string, error) {
	expireIn := viper.GetDuration("token.refresh_token_expires")
	if extendedRefresh {
		expireIn = viper.GetDuration("token.refresh_token_expires_extended")
	}
	refreshTokenClaims := jwt.MapClaims{
		"iss":     issuer,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(expireIn).Unix(),
		"user_id": userID,
	}
	return jwt.NewWithClaims(jwt.SigningMethodRS256, refreshTokenClaims).SignedString(privateKey)
}
