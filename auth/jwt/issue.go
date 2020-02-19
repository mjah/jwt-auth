package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mjah/jwt-auth/errors"
	"github.com/spf13/viper"
)

// IssueAccessToken issues an access token.
func (atc *AccessTokenClaims) IssueAccessToken() (string, *errors.ErrorCode) {
	tokenString, err := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		jwt.MapClaims{
			"type":    "access",
			"iss":     viper.GetString("token.issuer"),
			"iat":     atc.Iat,
			"exp":     atc.Exp,
			"user_id": atc.UserID,
			"role":    atc.Role,
		},
	).SignedString(privateKey)

	if err != nil {
		return "", errors.New(errors.AccessTokenIssueFailed, err)
	}

	return tokenString, nil
}

// IssueRefreshToken issues a refresh token.
func (rtc *RefreshTokenClaims) IssueRefreshToken() (string, *errors.ErrorCode) {
	tokenString, err := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		jwt.MapClaims{
			"type":    "refresh",
			"iss":     viper.GetString("token.issuer"),
			"iat":     rtc.Iat,
			"exp":     rtc.Exp,
			"user_id": rtc.UserID,
		},
	).SignedString(privateKey)

	if err != nil {
		return "", errors.New(errors.RefreshTokenIssueFailed, err)
	}

	return tokenString, nil
}
