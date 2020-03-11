package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/auth"
	"github.com/mjah/jwt-auth/auth/jwt"
)

// RefreshToken route handler.
func RefreshToken(c *gin.Context) {
	var details auth.RefreshTokenDetails

	claims := c.MustGet("refresh_token_claims").(jwt.RefreshTokenClaims)

	details.UserID = claims.UserID

	accessToken, errCode := details.RefreshToken()
	if errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"message": errCode.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Access token refreshed.",
		"access_token": accessToken,
	})
}
