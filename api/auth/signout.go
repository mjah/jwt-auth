package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/auth"
	"github.com/mjah/jwt-auth/auth/jwt"
	"github.com/spf13/viper"
)

// SignOut route handler.
func SignOut(c *gin.Context) {
	var details auth.SignOutDetails

	details.TokenString = c.MustGet("refresh_token_string").(string)
	details.Claims = c.MustGet("refresh_token_claims").(jwt.RefreshTokenClaims)

	if errCode := details.SignOut(); errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"error": errCode.OmitDetailsInProd()})
		return
	}

	if viper.GetBool("token.access_token.transport.cookies") {
		c.SetCookie("access_token", "", -1, "/", "", viper.GetBool("serve.cookies_secure"), true)
	}

	if viper.GetBool("token.refresh_token.transport.cookies") {
		c.SetCookie("refresh_token", "", -1, "/", "", viper.GetBool("serve.cookies_secure"), true)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Signed Out.",
	})
}

// SignOutAll route handler.
func SignOutAll(c *gin.Context) {
	var details auth.SignOutDetails

	details.Claims = c.MustGet("refresh_token_claims").(jwt.RefreshTokenClaims)

	if errCode := details.SignOutAll(); errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"error": errCode.OmitDetailsInProd()})
		return
	}

	if viper.GetBool("token.access_token.transport.cookies") {
		c.SetCookie("access_token", "", -1, "/", "", viper.GetBool("serve.cookies_secure"), true)
	}

	if viper.GetBool("token.refresh_token.transport.cookies") {
		c.SetCookie("refresh_token", "", -1, "/", "", viper.GetBool("serve.cookies_secure"), true)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Signed Out All.",
	})
}
