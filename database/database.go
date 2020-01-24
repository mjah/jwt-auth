package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // required
	"github.com/mjah/jwt-auth/logger"
	"github.com/spf13/viper"
)

// Connect ...
func Connect() *gorm.DB {
	environment := viper.GetString("environment")
	host := viper.GetString("postgres.host")
	port := viper.GetString("postgres.port")
	username := viper.GetString("postgres.username")
	password := viper.GetString("postgres.password")
	database := viper.GetString("postgres.database")

	dbDetails := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, username, password, database)

	db, err := gorm.Open("postgres", dbDetails)
	if err != nil {
		logger.Log().Fatal("Failed to connect to database.")
	}

	db.SingularTable(true)

	if environment != "production" {
		db.LogMode(true)
	}

	return db
}
