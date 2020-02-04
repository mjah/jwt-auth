package jwt

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return publicKey, nil
}

// ValidateToken ...
func ValidateToken(tokenString string) (*jwt.Token, error) {
	// to-do: check against database if revoked
	return jwt.Parse(tokenString, keyFunc) // to-do: return errorcode
}
