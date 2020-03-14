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

// ValidateToken checks if a token is valid.
func ValidateToken(tokenString string) (*jwt.Token, *errors.ErrorCode) {
	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return nil, errors.New(errors.JWTTokenInvalid, err.Error())
	}
	return token, nil
}
