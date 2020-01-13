package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Confirm ...
func Confirm(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
