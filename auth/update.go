package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Update ...
func Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
