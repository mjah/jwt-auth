package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User details.
type User struct {
	gorm.Model
	RoleID                   uint   `gorm:"not null"`
	Email                    string `gorm:"type:varchar(128);unique_index;not null"`
	Username                 string `gorm:"type:varchar(40);unique_index;not null"`
	Password                 string `gorm:"type:varchar(60);not null"`
	FirstName                string `gorm:"type:varchar(32);not null"`
	LastName                 string `gorm:"type:varchar(32);not null"`
	IsConfirmedEmail         bool   `gorm:"not null;default:false"`
	IsActive                 bool   `gorm:"not null;default:true"`
	ConfirmEmailToken        string
	ConfirmEmailTokenExpires time.Time
	ResetPassToken           string
	ResetPassTokenExpires    time.Time
	LastSignin               time.Time
	FailedSignin             time.Time
}

// Role details.
type Role struct {
	gorm.Model
	Role        string `gorm:"type:varchar(32);unique_index;not null"`
	Permissions string
}

// TokenRevocation details.
type TokenRevocation struct {
	gorm.Model
	UserID          uint `gorm:"not null"`
	RefreshToken    string
	RevokeAllBefore time.Time
}
