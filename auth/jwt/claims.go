package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mjah/jwt-auth/errors"
)

// AccessTokenClaims ...
type AccessTokenClaims struct {
	Iat    int64
	Exp    int64
	UserID uint
	Role   string
}

// RefreshTokenClaims ...
type RefreshTokenClaims struct {
	Iat    int64
	Exp    int64
	UserID uint
}

// ParseAccessTokenClaims ...
func ParseAccessTokenClaims(token *jwt.Token) (AccessTokenClaims, *errors.ErrorCode) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return AccessTokenClaims{}, errors.New(errors.AccessTokenClaimsParseFailed, nil)
	}

	if claims["type"].(string) != "access" {
		return AccessTokenClaims{}, errors.New(errors.AccessTokenClaimsParseFailed, nil)
	}

	atc := AccessTokenClaims{
		Iat:    int64(claims["iat"].(float64)),
		Exp:    int64(claims["exp"].(float64)),
		UserID: uint(claims["user_id"].(float64)),
		Role:   claims["role"].(string),
	}
	return atc, nil
}

// ParseRefreshTokenClaims ...
func ParseRefreshTokenClaims(token *jwt.Token) (RefreshTokenClaims, *errors.ErrorCode) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return RefreshTokenClaims{}, errors.New(errors.RefreshTokenClaimsParseFailed, nil)
	}

	if claims["type"].(string) != "refresh" {
		return RefreshTokenClaims{}, errors.New(errors.RefreshTokenClaimsParseFailed, nil)
	}

	rtc := RefreshTokenClaims{
		Iat:    int64(claims["iat"].(float64)),
		Exp:    int64(claims["exp"].(float64)),
		UserID: uint(claims["user_id"].(float64)),
	}
	return rtc, nil
}
