package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkRevoked() {

}

// RefreshToken ...
func RefreshToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
