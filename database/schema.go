package database

import "github.com/jinzhu/gorm"

// User ...
type User struct {
	gorm.Model
	RoleID                int
	Email                 string `gorm:"type:varchar(128);unique_index"`
	Username              string
	PasswordHash          string
	Password              string
	FirstName             string
	LastName              string
	IsConfirmed           string
	IsActive              string
	ResetPassToken        string
	ResetPassTokenCreated string
	LastLogin             string
	FailedLogin           string
	LockExpires           string
}

// Role ...
type Role struct {
	gorm.Model
	Role        string `gorm:"type:varchar(32);unique_index"`
	Permissions string
}

// TokenRevocation ...
type TokenRevocation struct {
	gorm.Model
	UserID        int
	RefreshToken  string
	LogoutAll     bool
	LogoutAllTime string
}

// Migrate ...
func Migrate() {
	db := Connect()
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Role{})
	db.AutoMigrate(&TokenRevocation{})
}
