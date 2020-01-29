package queue

import (
	"fmt"

	"github.com/mjah/jwt-auth/logger"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var amqpInstance *amqp.Connection

// Connect ...
func Connect() (*amqp.Connection, error) {
	amqpDetails := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		viper.GetString("amqp.username"),
		viper.GetString("amqp.password"),
		viper.GetString("amqp.host"),
		viper.GetString("amqp.port"),
	)

	conn, err := amqp.Dial(amqpDetails)
	if err != nil {
		logger.Log().Fatal("Failed to connect to message-broker.")
		return nil, err
	}

	amqpInstance = conn

	return conn, nil
}

// GetConnection ...
func GetConnection() (*amqp.Connection, error) {
	return amqpInstance, nil
}
