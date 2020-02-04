package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/auth"
	"github.com/mjah/jwt-auth/auth/jwt"
)

// SignOut ...
func SignOut(c *gin.Context) {
	refreshTokenString := c.MustGet("refresh_token_string").(string)
	refreshTokenClaims := c.MustGet("refresh_token_claims").(jwt.RefreshTokenClaims)

	if errCode := auth.SignOut(refreshTokenString, refreshTokenClaims); errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"message": errCode.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Signed Out.",
	})
}

// SignOutAll ...
func SignOutAll(c *gin.Context) {
	refreshTokenClaims := c.MustGet("refresh_token_claims").(jwt.RefreshTokenClaims)

	if errCode := auth.SignOutAll(refreshTokenClaims); errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"message": errCode.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Signed Out All.",
	})
}
