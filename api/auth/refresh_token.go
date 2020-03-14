package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/auth"
	"github.com/mjah/jwt-auth/auth/jwt"
	"github.com/spf13/viper"
)

// RefreshToken route handler.
func RefreshToken(c *gin.Context) {
	var details auth.RefreshTokenDetails

	claims := c.MustGet("refresh_token_claims").(jwt.RefreshTokenClaims)

	details.UserID = claims.UserID

	accessToken, errCode := details.RefreshToken()
	if errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"error": errCode.OmitDetailsInProd()})
		return
	}

	if viper.GetBool("token.access_token.transport.cookies") {
		maxAge := int(viper.GetDuration("token.access_token.expires").Seconds())
		c.SetCookie("access_token", accessToken, maxAge, "/", "", viper.GetBool("serve.cookies_secure"), true)
	}

	if viper.GetBool("token.access_token.transport.json_response") {
		c.JSON(http.StatusOK, gin.H{
			"message":      "Access token refreshed.",
			"access_token": accessToken,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Access token refreshed.",
		})
	}
}
