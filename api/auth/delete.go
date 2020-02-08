package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/auth"
	"github.com/mjah/jwt-auth/auth/jwt"
)

// Delete ...
func Delete(c *gin.Context) {
	refreshTokenClaims := c.MustGet("refresh_token_claims").(jwt.RefreshTokenClaims)

	if errCode := auth.Delete(refreshTokenClaims); errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"message": errCode.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Account deleted.",
	})
}
