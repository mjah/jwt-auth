package database

import (
	"github.com/spf13/viper"
)

// Migrate ...
func Migrate() error {
	db, err := GetConnection()
	if err != nil {
		return err
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

	return nil
}
