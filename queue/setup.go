package queue

import (
	"fmt"
	"time"

	"github.com/mjah/jwt-auth/logger"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

// Queue ...
type Queue struct {
	name        string
	contentType string

	connection *amqp.Connection
	notifyErr  chan *amqp.Error
	channel    *amqp.Channel
	consumers  []messageConsumer
}

// New ...
func New(queueName string) (*Queue, error) {
	q := &Queue{
		name:        queueName,
		contentType: "text/plain",
	}

	if err := q.setup(); err != nil {
		return nil, err
	}

	return q, nil
}

// Close ...
func (q *Queue) Close() error {
	if err := q.connection.Close(); err != nil {
		return err
	}
	return nil
}

func (q *Queue) setup() error {
	if err := q.openConnection(); err != nil {
		return err
	}

	if err := q.openChannel(); err != nil {
		return err
	}

	if err := q.setQos(); err != nil {
		return err
	}

	if err := q.declareQueue(); err != nil {
		return err
	}

	go q.handleError()

	return nil
}

func (q *Queue) openConnection() error {
	amqpDetails := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		viper.GetString("amqp.username"),
		viper.GetString("amqp.password"),
		viper.GetString("amqp.host"),
		viper.GetString("amqp.port"),
	)

	conn, err := amqp.Dial(amqpDetails)
	if err != nil {
		return err
	}
	q.connection = conn
	q.notifyErr = q.connection.NotifyClose(make(chan *amqp.Error))

	return nil
}

func (q *Queue) handleError() {
	err := <-q.notifyErr

	if err != nil {
		logger.Log().Error(err)

		retries := 0
		sleepSec := 0

		for {
			retries++
			if retries <= 60 {
				sleepSec++
			}

			logger.Log().Info("Attempting message-broker reconnection.")
			if err := q.setup(); err != nil {
				logger.Log().Error("Failed message-broker reconnection: ", err)
				time.Sleep(time.Duration(sleepSec) * time.Second)
				continue
			}
			logger.Log().Info("Reconnected to message-broker.")

			logger.Log().Info("Attempting consumer(s) recovery.")
			if err := q.recoverConsumer(); err != nil {
				logger.Log().Error("Failed consumer(s) recovery: ", err)
				time.Sleep(time.Duration(sleepSec) * time.Second)
				continue
			}
			logger.Log().Info("Recovered consumer(s).")

			return
		}
	}
}

func (q *Queue) openChannel() error {
	ch, err := q.connection.Channel()
	q.channel = ch
	return err
}

func (q *Queue) setQos() error {
	err := q.channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	return err
}

func (q *Queue) declareQueue() error {
	_, err := q.channel.QueueDeclare(
		q.name, // name
		true,   // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	return err
}
