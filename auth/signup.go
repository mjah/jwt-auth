package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignUp ...
func SignUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}