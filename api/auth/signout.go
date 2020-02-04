package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignOut ...
func SignOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// SignOutAll ...
func SignOutAll(c *gin.Context) {

}
