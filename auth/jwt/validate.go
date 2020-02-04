package jwt

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/mjah/jwt-auth/errors"
)

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return publicKey, nil
}

// ValidateToken ...
func ValidateToken(tokenString string) (*jwt.Token, *errors.ErrorCode) {
	// to-do: check against database if revoked
	//        check if locked
	//        check if active

	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return nil, errors.New(errors.TokenValidationFailed, err)
	}

	return token, nil
}
