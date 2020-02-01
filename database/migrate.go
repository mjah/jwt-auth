package database

import (
	"github.com/mjah/jwt-auth/logger"
	"github.com/spf13/viper"
)

// Migrate ...
func Migrate() {
	db, err := GetConnection()
	if err != nil {
		logger.Log().Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Role{})
	db.AutoMigrate(&TokenRevocation{})
	db.AutoMigrate(&EmailQueue{})

	for _, role := range viper.GetStringSlice("roles.define") {
		submitRole := &Role{Role: role}
		db.FirstOrCreate(&Role{}, submitRole)
	}
}
