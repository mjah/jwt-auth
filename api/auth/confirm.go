package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/auth"
	"github.com/mjah/jwt-auth/errors"
)

// Confirm ...
func Confirm(c *gin.Context) {
	var details auth.ConfirmDetails

	if err := c.BindJSON(&details); err != nil {
		err := errors.New(errors.ConfirmDetailsInvalid, err)
		c.AbortWithStatusJSON(err.HTTPStatus, gin.H{"message": err.OmitDetailsInProd()})
		return
	}

	if err := details.Confirm(); err != nil {
		c.AbortWithStatusJSON(err.HTTPStatus, gin.H{"message": err.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Account confirmed.",
	})
}

// SendConfirmEmail ...
func SendConfirmEmail(c *gin.Context) {

}
