package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignOut ...
func SignOut(c *gin.Context) {
	c.SetCookie("access_token", "", 0, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}