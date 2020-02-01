package api

import (
	"net/http"
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

// ValidateAccessTokenMiddleware ...
func ValidateAccessTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenBearer := c.GetHeader("Authorization")

		if len(tokenBearer) == 0 {
			err := errors.New(errors.AuthorizationBearerTokenEmpty, nil)
			c.AbortWithStatusJSON(err.HTTPStatus, gin.H{"message": err})
			return
		}

		tokenString := stripBearerPrefix(tokenBearer)
		if _, err := jwt.ValidateAccessToken(tokenString); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// claims, ok := token.Claims.(jwt.MapClaims)
		// if !ok || !token.Valid {
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// 	return
		// }
	}
}
