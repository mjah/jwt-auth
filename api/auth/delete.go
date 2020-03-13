package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/auth"
	"github.com/mjah/jwt-auth/auth/jwt"
)

// Delete route handler.
func Delete(c *gin.Context) {
	var details auth.DeleteDetails

	claims := c.MustGet("refresh_token_claims").(jwt.RefreshTokenClaims)

	details.UserID = claims.UserID

	if errCode := details.Delete(); errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"error": errCode.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Account deleted.",
	})
}
