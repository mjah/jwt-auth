package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/auth"
	"github.com/mjah/jwt-auth/errors"
)

// SignIn ...
func SignIn(c *gin.Context) {
	var details auth.SignInDetails

	if err := c.BindJSON(&details); err != nil {
		errCode := errors.New(errors.SignInDetailsInvalid, err)
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"message": errCode.OmitDetailsInProd()})
		return
	}

	accessToken, refreshToken, errCode := details.SignIn()
	if errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"message": errCode.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
