package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/auth"
	"github.com/mjah/jwt-auth/auth/jwt"
)

// RefreshToken ...
func RefreshToken(c *gin.Context) {
	refreshTokenClaims := c.MustGet("refresh_token_claims").(jwt.RefreshTokenClaims)

	accessToken, errCode := auth.RefreshToken(refreshTokenClaims)
	if errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"message": errCode.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Access token refreshed.",
		"access_token": accessToken,
	})
}
