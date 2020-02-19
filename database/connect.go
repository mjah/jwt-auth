// Package database implements postgres connection, schema, and migration.
package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // required
	"github.com/mjah/jwt-auth/logger"
	"github.com/spf13/viper"
)

var dbInstance *gorm.DB

// Connect opens a database connection and returns it.
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		viper.GetString("postgres.host"),
		viper.GetString("postgres.port"),
		viper.GetString("postgres.username"),
		viper.GetString("postgres.password"),
		viper.GetString("postgres.database"),
	))
	if err != nil {
		return nil, err
	}

	if viper.GetString("environment") == "production" {
		db.LogMode(false)
	} else {
		db.LogMode(true)
	}

	db.SingularTable(true)
	// db.SetLogger(logger.Log())

	dbInstance = db

	return db, nil
}

// GetConnection returns an existing connection, otherwise opens a new one.
func GetConnection() (*gorm.DB, error) {
	if err := dbInstance.DB().Ping(); err != nil {
		if _, err := Connect(); err != nil {
			logger.Log().Error(err)
			return nil, err
		}
	}
	return dbInstance, nil
}
