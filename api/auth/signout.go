package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/auth"
	"github.com/mjah/jwt-auth/auth/jwt"
)

// SignOut ...
func SignOut(c *gin.Context) {
	var details auth.SignOutDetails

	details.TokenString = c.MustGet("refresh_token_string").(string)
	details.Claims = c.MustGet("refresh_token_claims").(jwt.RefreshTokenClaims)

	if errCode := details.SignOut(); errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"message": errCode.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Signed Out.",
	})
}

// SignOutAll ...
func SignOutAll(c *gin.Context) {
	var details auth.SignOutDetails

	details.Claims = c.MustGet("refresh_token_claims").(jwt.RefreshTokenClaims)

	if errCode := details.SignOutAll(); errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"message": errCode.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Signed Out All.",
	})
}
