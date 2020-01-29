package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	gorm.Model
	RoleID                uint   `gorm:"not null"`
	Email                 string `gorm:"type:varchar(128);unique_index;not null"`
	Username              string `gorm:"type:varchar(40);unique_index;not null"`
	Password              string `gorm:"type:varchar(60);not null"`
	FirstName             string `gorm:"type:varchar(32);not null"`
	LastName              string `gorm:"type:varchar(32);not null"`
	IsConfirmed           bool   `gorm:"not null;default:false"`
	IsActive              bool   `gorm:"not null;default:false"`
	ConfirmToken          string
	ConfirmTokenExpires   time.Time
	ResetPassToken        string
	ResetPassTokenExpires time.Time
	LastLogin             time.Time
	FailedLogin           time.Time
	LockExpires           time.Time
}

// Role ...
type Role struct {
	gorm.Model
	Role        string `gorm:"type:varchar(32);unique_index;not null"`
	Permissions string
}

// TokenRevocation ...
type TokenRevocation struct {
	gorm.Model
	UserID       uint `gorm:"not null"`
	RefreshToken string
	LogoutAll    time.Time
}

// EmailQueue ...
type EmailQueue struct {
	gorm.Model
	UserID            uint   `gorm:"not null"`
	RecipientEmail    string `gorm:"type:varchar(128);not null"`
	RecipientName     string `gorm:"type:varchar(32);not null"`
	EmailType         string `gorm:"type:varchar(32);not null"`
	MessageParameters string
	ProcessedAt       time.Time
}
