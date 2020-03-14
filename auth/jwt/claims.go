package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mjah/jwt-auth/errors"
)

// AccessTokenClaims holds the access token claims details.
type AccessTokenClaims struct {
	Iat    int64
	Exp    int64
	UserID uint
	Role   string
}

// RefreshTokenClaims holds the refresh token claims details.
type RefreshTokenClaims struct {
	Iat    int64
	Exp    int64
	UserID uint
}

// ParseAccessTokenClaims parses the access token claims.
func ParseAccessTokenClaims(token *jwt.Token) (AccessTokenClaims, *errors.ErrorCode) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return AccessTokenClaims{}, errors.New(errors.JWTTokenInvalid, "")
	}

	if claims["type"].(string) != "access" {
		return AccessTokenClaims{}, errors.New(errors.JWTTokenInvalid, "")
	}

	atc := AccessTokenClaims{
		Iat:    int64(claims["iat"].(float64)),
		Exp:    int64(claims["exp"].(float64)),
		UserID: uint(claims["user_id"].(float64)),
		Role:   claims["role"].(string),
	}
	return atc, nil
}

// ParseRefreshTokenClaims parses the refresh token claims.
func ParseRefreshTokenClaims(token *jwt.Token) (RefreshTokenClaims, *errors.ErrorCode) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return RefreshTokenClaims{}, errors.New(errors.JWTTokenInvalid, "")
	}

	if claims["type"].(string) != "refresh" {
		return RefreshTokenClaims{}, errors.New(errors.JWTTokenInvalid, "")
	}

	rtc := RefreshTokenClaims{
		Iat:    int64(claims["iat"].(float64)),
		Exp:    int64(claims["exp"].(float64)),
		UserID: uint(claims["user_id"].(float64)),
	}
	return rtc, nil
}
