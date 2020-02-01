package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const issuer = "jwt-auth-server"
const expireIn = time.Hour * 24

// IssueAccessToken ...
func IssueAccessToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": issuer,
		"exp": time.Now().Add(expireIn).Unix(),
		"grp": "admin",
	})
	return token.SignedString(privateKey)
}

// IssueRefreshToken ...
func IssueRefreshToken(extendedRefresh bool) (string, error) {
	// if extendedRefresh {
	// 	expireIn := viper
	// } else {
	// 	expireIn :=
	// }
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": issuer,
		"exp": time.Now().Add(expireIn).Unix(),
		"grp": "admin",
	})
	return token.SignedString(privateKey)
}
