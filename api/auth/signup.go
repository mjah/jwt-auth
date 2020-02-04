package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/auth"
	"github.com/mjah/jwt-auth/errors"
)

// SignUp ...
func SignUp(c *gin.Context) {
	var details auth.SignUpDetails

	if err := c.BindJSON(&details); err != nil {
		err := errors.New(errors.SignUpDetailsInvalid, err)
		c.AbortWithStatusJSON(err.HTTPStatus, gin.H{"message": err.OmitDetailsInProd()})
		return
	}

	if err := details.SignUp(); err != nil {
		c.AbortWithStatusJSON(err.HTTPStatus, gin.H{"message": err.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Account created.",
	})
}
