package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/database"
)

// SignUpDetails ...
type SignUpDetails struct {
	Email     string `json:"email" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

func sendConfirmationEmail() {

}

// CreateUser ...
func CreateUser(user *SignUpDetails) error {
	db := database.Connect()
	defer db.Close()

	role := &database.Role{Role: "Guest"}
	db.Where("role = ?", "Guest").First(&role)

	submitUser := &database.User{
		RoleID:    role.ID,
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	err := db.FirstOrCreate(&database.User{}, submitUser).Error
	if err != nil {
		return err
	}

	return nil
}

// SignUp ...
func SignUp(c *gin.Context) {
	var user SignUpDetails

	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = CreateUser(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Account created.",
	})
}
