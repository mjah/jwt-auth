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
	ResetPassToken        string
	ResetPassTokenCreated *time.Time
	LastLogin             *time.Time
	FailedLogin           *time.Time
	LockExpires           *time.Time
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
	LogoutAll    *time.Time
}

// EmailQueue ...
type EmailQueue struct {
	gorm.Model
	UserID            uint   `gorm:"not null"`
	RecipientEmail    string `gorm:"type:varchar(128);not null"`
	RecipientName     string `gorm:"type:varchar(32);not null"`
	EmailType         string `gorm:"type:varchar(32);not null"`
	MessageParameters string
	ProcessedAt       *time.Time
}

// Migrate ...
func Migrate() {
	db := GetConnection()
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Role{})
	db.AutoMigrate(&TokenRevocation{})
	db.AutoMigrate(&EmailQueue{})

	adminRole := &Role{Role: "Admin"}
	db.FirstOrCreate(adminRole, adminRole)

	memberRole := &Role{Role: "Member"}
	db.FirstOrCreate(memberRole, memberRole)

	guestRole := &Role{Role: "Guest"}
	db.FirstOrCreate(guestRole, guestRole)
}
