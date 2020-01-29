package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/utils"
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
	db := database.GetConnection()
	defer db.Close()

	role := &database.Role{Role: "Guest"}
	db.Where("role = ?", "Guest").First(&role)

	generatedPassword, err := utils.GeneratePassword(user.Password)
	if err != nil {
		return err
	}

	submitUser := &database.User{
		RoleID:    role.ID,
		Email:     user.Email,
		Username:  user.Username,
		Password:  generatedPassword,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	err = db.FirstOrCreate(&database.User{}, submitUser).Error
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
