package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/auth"
)

// SignIn ...
func SignIn(c *gin.Context) {
	var details auth.SignInDetails

	if err := c.BindJSON(&details); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tokenString, err := details.SignIn()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  tokenString,
		"token_type":    "bearer",
		"refresh_token": "00000000000",
	})
}
