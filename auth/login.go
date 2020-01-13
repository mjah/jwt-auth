package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login ...
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
