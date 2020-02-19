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
		errCode := errors.New(errors.DetailsInvalid, err)
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"message": errCode.OmitDetailsInProd()})
		return
	}

	if errCode := details.Confirm(); errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"message": errCode.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Account confirmed.",
	})
}

// SendConfirmEmail ...
func SendConfirmEmail(c *gin.Context) {
	var details auth.SendConfirmEmailDetails

	if err := c.BindJSON(&details); err != nil {
		errCode := errors.New(errors.DetailsInvalid, err)
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"message": errCode.OmitDetailsInProd()})
		return
	}

	if errCode := details.SendConfirmEmail(); errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"message": errCode.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sent confirm email.",
	})
}
