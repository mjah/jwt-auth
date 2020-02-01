package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const issuer = "auth-server"
const expireIn = time.Hour * 24

// ValidateAccessToken ...
func ValidateAccessToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
}

// ValidateRefreshToken ...
func ValidateRefreshToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
}

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
func IssueRefreshToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": issuer,
		"exp": time.Now().Add(expireIn).Unix(),
		"grp": "admin",
	})
	return token.SignedString(privateKey)
}

// RevokeRefreshToken ...
func RevokeRefreshToken() {

}
