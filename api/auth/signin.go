package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/auth"
	"github.com/mjah/jwt-auth/errors"
	"github.com/spf13/viper"
)

// SignIn route handler.
func SignIn(c *gin.Context) {
	var details auth.SignInDetails

	if err := c.BindJSON(&details); err != nil {
		errCode := errors.New(errors.DetailsInvalid, err.Error())
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"error": errCode.OmitDetailsInProd()})
		return
	}

	accessToken, refreshToken, errCode := details.SignIn()
	if errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"error": errCode.OmitDetailsInProd()})
		return
	}

	if viper.GetBool("token.access_token.transport.cookies") {
		maxAge := int(viper.GetDuration("token.access_token.expires").Seconds())
		c.SetCookie("access_token", accessToken, maxAge, "/", "", viper.GetBool("serve.cookies_secure"), true)
	}

	if viper.GetBool("token.refresh_token.transport.cookies") {
		maxAge := int(viper.GetDuration("token.refresh_token.expires").Seconds())
		if details.RememberMe {
			maxAge = int(viper.GetDuration("token.refresh_token.expires_extended").Seconds())
		}
		c.SetCookie("refresh_token", refreshToken, maxAge, "/", "", viper.GetBool("serve.cookies_secure"), true)
	}

	// to-do: clean this up
	if viper.GetBool("token.access_token.transport.json_response") &&
		viper.GetBool("token.refresh_token.transport.json_response") {
		c.JSON(http.StatusOK, gin.H{
			"message":       "Signed In.",
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	} else if viper.GetBool("token.access_token.transport.json_response") {
		c.JSON(http.StatusOK, gin.H{
			"message":      "Signed In.",
			"access_token": accessToken,
		})
	} else if viper.GetBool("token.refresh_token.transport.json_response") {
		c.JSON(http.StatusOK, gin.H{
			"message":       "Signed In.",
			"refresh_token": refreshToken,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Signed In.",
		})
	}
}
