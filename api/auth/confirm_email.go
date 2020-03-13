package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/auth"
	"github.com/mjah/jwt-auth/errors"
)

// ConfirmEmail route handler.
func ConfirmEmail(c *gin.Context) {
	var details auth.ConfirmEmailDetails

	if err := c.BindJSON(&details); err != nil {
		errCode := errors.New(errors.DetailsInvalid, err)
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"error": errCode.OmitDetailsInProd()})
		return
	}

	if errCode := details.ConfirmEmail(); errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"error": errCode.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Email confirmed.",
	})
}

// SendConfirmEmail route handler.
func SendConfirmEmail(c *gin.Context) {
	var details auth.SendConfirmEmailDetails

	if err := c.BindJSON(&details); err != nil {
		errCode := errors.New(errors.DetailsInvalid, err)
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"error": errCode.OmitDetailsInProd()})
		return
	}

	if errCode := details.SendConfirmEmail(); errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"error": errCode.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sent confirm email.",
	})
}
