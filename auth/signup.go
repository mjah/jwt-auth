package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userDetails struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

func sendConfirmationEmail() {

}

// SignUp ...
func SignUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
