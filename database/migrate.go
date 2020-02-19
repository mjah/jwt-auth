package database

import (
	"github.com/spf13/viper"
)

// Migrate automatically migrates the database to the latest schema.
func Migrate() error {
	db, err := GetConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Role{})
	db.AutoMigrate(&TokenRevocation{})

	for _, role := range viper.GetStringSlice("roles.define") {
		submitRole := &Role{Role: role}
		db.FirstOrCreate(&Role{}, submitRole)
	}

	return nil
}
