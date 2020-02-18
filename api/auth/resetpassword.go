package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/auth"
	"github.com/mjah/jwt-auth/errors"
)

// ResetPassword ...
func ResetPassword(c *gin.Context) {
	var details auth.ResetPasswordDetails

	if err := c.BindJSON(&details); err != nil {
		errCode := errors.New(errors.ResetPasswordDetailsInvalid, err)
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"message": errCode.OmitDetailsInProd()})
		return
	}

	if errCode := details.ResetPassword(); errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"message": errCode.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Account password reset.",
	})
}

// SendResetPasswordEmail ...
func SendResetPasswordEmail(c *gin.Context) {
	var details auth.SendResetPasswordEmailDetails

	if err := c.BindJSON(&details); err != nil {
		errCode := errors.New(errors.SendResetPasswordEmailDetailsInvalid, err)
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"message": errCode.OmitDetailsInProd()})
		return
	}

	if errCode := details.SendResetPasswordEmail(); errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPStatus, gin.H{"message": errCode.OmitDetailsInProd()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sent reset password email.",
	})
}
