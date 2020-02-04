package api

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/auth/jwt"
	"github.com/mjah/jwt-auth/errors"
)

func stripBearerPrefix(tokenBearer string) string {
	if len(tokenBearer) > 6 && strings.ToUpper(tokenBearer[0:7]) == "BEARER " {
		return tokenBearer[7:]
	}
	return tokenBearer
}

// ValidateRefreshTokenMiddleware ...
func ValidateRefreshTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenBearer := c.GetHeader("Authorization")

		if len(tokenBearer) == 0 {
			err := errors.New(errors.AuthorizationBearerTokenEmpty, nil)
			c.AbortWithStatusJSON(err.HTTPStatus, gin.H{"message": err})
			return
		}

		tokenString := stripBearerPrefix(tokenBearer)
		token, err := jwt.ValidateToken(tokenString)
		if err != nil {
			err := errors.New(errors.RefreshTokenValidationFailed, err)
			c.AbortWithStatusJSON(err.HTTPStatus, gin.H{"message": err})
			return
		}

		claims, err := jwt.ParseRefreshTokenClaims(token)
		if err != nil {
			err := errors.New(errors.RefreshTokenClaimsParseFailed, err)
			c.AbortWithStatusJSON(err.HTTPStatus, gin.H{"message": err})
			return
		}

		c.Set("refresh_token_string", tokenString)
		c.Set("refresh_token_claims", claims)
	}
}
