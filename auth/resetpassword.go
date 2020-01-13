package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResetPassword ...
func ResetPassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
