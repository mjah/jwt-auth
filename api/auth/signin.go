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
		err := errors.New(errors.SignInDetailsInvalid, err)
		c.AbortWithStatusJSON(err.HTTPStatus, gin.H{"message": err.OmitDetailsInProd()})
		return
	}

	accessToken, refreshToken, err := details.SignIn()
	if err != nil {
		c.AbortWithStatusJSON(err.HTTPStatus, gin.H{"message": err.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
