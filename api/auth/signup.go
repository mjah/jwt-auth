package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/auth"
)

// SignUp ...
func SignUp(c *gin.Context) {
	var user auth.SignUpDetails

	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = user.SignUp()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Account created.",
	})
}
