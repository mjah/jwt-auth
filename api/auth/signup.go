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
		errCode := errors.New(errors.DetailsInvalid, err)
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"message": errCode.OmitDetailsInProd()})
		return
	}

	if errCode := details.SignUp(); errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"message": errCode.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Account created.",
	})
}
