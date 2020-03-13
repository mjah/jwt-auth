package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/auth"
	"github.com/mjah/jwt-auth/auth/jwt"
)

// UserDetails route handler.
func UserDetails(c *gin.Context) {
	var details auth.UserDetailsDetails

	claims := c.MustGet("refresh_token_claims").(jwt.RefreshTokenClaims)

	details.UserID = claims.UserID

	userDetails, errCode := details.UserDetails()
	if errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"error": errCode.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Retrieved account details.",
		"details": userDetails,
	})
}
