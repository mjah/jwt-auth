package api

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/auth/jwt"
	"github.com/mjah/jwt-auth/errors"
	"github.com/spf13/viper"
)

func stripBearerPrefix(tokenBearer string) string {
	if len(tokenBearer) > 6 && strings.ToUpper(tokenBearer[0:7]) == "BEARER " {
		return tokenBearer[7:]
	}
	return tokenBearer
}

// ValidateRefreshTokenMiddleware checks to see if the refresh token is valid
// before accessing private resources.
func ValidateRefreshTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string

		if viper.GetBool("token.refresh_token.transport.cookies") {
			var err error
			tokenString, err = c.Cookie("refresh_token")
			if err != nil {
				errCode := errors.New(errors.RefreshTokenCookieEmpty, err.Error())
				c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"error": errCode.OmitDetailsInProd()})
				return
			}
		} else {
			tokenBearer := c.GetHeader("Authorization")
			if len(tokenBearer) == 0 {
				errCode := errors.New(errors.AuthorizationBearerTokenEmpty, "")
				c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"error": errCode.OmitDetailsInProd()})
				return
			}
			tokenString = stripBearerPrefix(tokenBearer)
		}

		token, errCode := jwt.ValidateToken(tokenString)
		if errCode != nil {
			c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"error": errCode.OmitDetailsInProd()})
			return
		}

		claims, errCode := jwt.ParseRefreshTokenClaims(token)
		if errCode != nil {
			c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"error": errCode.OmitDetailsInProd()})
			return
		}

		if errCode := jwt.CheckRefreshTokenRevoked(claims.UserID, claims.Iat, tokenString); errCode != nil {
			c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"error": errCode.OmitDetailsInProd()})
			return
		}

		c.Set("refresh_token_string", tokenString)
		c.Set("refresh_token_claims", claims)
	}
}
