package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userCredentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func authenticateUserCredentials(user *userCredentials) bool {
	// authentication logic here

	// testing
	if user.Username == "moejay" && user.Password == "password123" {
		return true
	}
	return false
}

// Login ...
func Login(c *gin.Context) {
	var user userCredentials

	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	validUser := authenticateUserCredentials(&user)
	if validUser == false {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString, err := issueToken()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.SetCookie("access_token", tokenString, 86400, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  tokenString,
		"token_type":    "bearer",
		"refresh_token": "00000000000",
	})
}
