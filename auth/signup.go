package auth

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/utils"
)

// SignUpDetails ...
type SignUpDetails struct {
	Email     string `json:"email" binding:"required" valid:"email"`
	Username  string `json:"username" binding:"required" valid:"length(3|40)"`
	Password  string `json:"password" binding:"required" valid:"length(8|60)"`
	FirstName string `json:"first_name" binding:"required" valid:"length(1|32)"`
	LastName  string `json:"last_name" binding:"required" valid:"length(1|32)"`
}

// SignUp ...
func (details *SignUpDetails) SignUp() error {
	if _, err := govalidator.ValidateStruct(details); err != nil {
		return err
	}

	db, err := database.GetConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	role := &database.Role{Role: "Guest"}
	db.Where("role = ?", "Guest").First(&role)

	generatedPassword, err := utils.GeneratePassword(details.Password)
	if err != nil {
		return err
	}

	submitUser := &database.User{
		RoleID:              role.ID,
		Email:               details.Email,
		Username:            details.Username,
		Password:            generatedPassword,
		FirstName:           details.FirstName,
		LastName:            details.LastName,
		ConfirmToken:        utils.GenerateUUID(),
		ConfirmTokenExpires: time.Now().Add(time.Hour * 24),
	}

	err = db.FirstOrCreate(&database.User{}, submitUser).Error
	if err != nil {
		return err
	}

	return nil
}
