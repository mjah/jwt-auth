package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Delete ...
func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
