package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

const issuer = "jwt-auth-server"

var expireIn time.Duration

// IssueAccessToken ...
func IssueAccessToken() (string, error) {
	expireIn = viper.GetDuration("token.access_token_expires")
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": issuer,
		"exp": time.Now().Add(expireIn).Unix(),
		"grp": "admin",
	})
	return token.SignedString(privateKey)
}

// IssueRefreshToken ...
func IssueRefreshToken(extendedRefresh bool) (string, error) {
	if extendedRefresh {
		expireIn = viper.GetDuration("token.refresh_token_expires_extended")
	} else {
		expireIn = viper.GetDuration("token.refresh_token_expires")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": issuer,
		"exp": time.Now().Add(expireIn).Unix(),
		"grp": "admin",
	})
	return token.SignedString(privateKey)
}
