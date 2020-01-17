package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

// Connect ...
func Connect() *gorm.DB {
	host := viper.GetString("postgres.host")
	port := viper.GetString("postgres.port")
	username := viper.GetString("postgres.username")
	password := viper.GetString("postgres.password")
	database := viper.GetString("postgres.database")

	dbDetails := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, username, password, database)

	db, err := gorm.Open("postgres", dbDetails)
	if err != nil {
		panic("Failed to connect to database.")
	}

	db.SingularTable(true)

	return db
}
