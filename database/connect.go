package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // required
	"github.com/spf13/viper"
)

var dbInstance *gorm.DB

// Connect ...
func Connect() (*gorm.DB, error) {
	dbDetails := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		viper.GetString("postgres.host"),
		viper.GetString("postgres.port"),
		viper.GetString("postgres.username"),
		viper.GetString("postgres.password"),
		viper.GetString("postgres.database"),
	)

	db, err := gorm.Open("postgres", dbDetails)
	if err != nil {
		return nil, err
	}

	if viper.GetString("environment") != "production" {
		db.LogMode(true)
	}

	db.SingularTable(true)
	// db.SetLogger(logger.Log())

	dbInstance = db

	return db, nil
}

// GetConnection ...
func GetConnection() (*gorm.DB, error) {
	if err := dbInstance.DB().Ping(); err != nil {
		if _, err := Connect(); err != nil {
			return nil, err
		}
	}
	return dbInstance, nil
}
