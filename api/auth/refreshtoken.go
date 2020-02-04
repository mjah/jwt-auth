package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RefreshToken ...
func RefreshToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
