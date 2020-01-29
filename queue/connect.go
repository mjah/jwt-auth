package queue

import (
	"fmt"

	"github.com/mjah/jwt-auth/logger"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var amqpInstance *amqp.Connection

// Connect ...
func Connect() *amqp.Connection {
	host := viper.GetString("amqp.host")
	port := viper.GetString("amqp.port")
	username := viper.GetString("amqp.username")
	password := viper.GetString("amqp.password")

	amqpDetails := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port)

	conn, err := amqp.Dial(amqpDetails)
	if err != nil {
		logger.Log().Fatal("Failed to connect to message-broker.")
	}

	amqpInstance = conn

	return conn
}

// GetConnection ...
func GetConnection() *amqp.Connection {
	return amqpInstance
}
