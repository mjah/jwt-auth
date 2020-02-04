package jwt

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

// AccessTokenClaims ...
type AccessTokenClaims struct {
	Iss    string
	Iat    int64
	Exp    int64
	UserID uint
	Role   string
}

// RefreshTokenClaims ...
type RefreshTokenClaims struct {
	Iss    string
	Iat    int64
	Exp    int64
	UserID uint
}

// ParseAccessTokenClaims ...
func ParseAccessTokenClaims(token *jwt.Token) (AccessTokenClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return AccessTokenClaims{}, errors.New("could not parse claims")
	}
	if claims["sub"] != "access" {
		return AccessTokenClaims{}, errors.New("not an access token")
	}
	atc := AccessTokenClaims{
		Iss:    claims["iss"].(string),
		Iat:    int64(claims["iat"].(float64)),
		Exp:    int64(claims["exp"].(float64)),
		UserID: uint(claims["user_id"].(float64)),
		Role:   claims["role"].(string),
	}
	return atc, nil
}

// ParseRefreshTokenClaims ...
func ParseRefreshTokenClaims(token *jwt.Token) (RefreshTokenClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return RefreshTokenClaims{}, errors.New("could not parse claims")
	}
	if claims["sub"] != "refresh" {
		return RefreshTokenClaims{}, errors.New("not a refresh token")
	}
	rtc := RefreshTokenClaims{
		Iss:    claims["iss"].(string),
		Iat:    int64(claims["iat"].(float64)),
		Exp:    int64(claims["exp"].(float64)),
		UserID: uint(claims["user_id"].(float64)),
	}
	return rtc, nil
}
