package queue

import "github.com/streadway/amqp"

// SetContentType ...
func (q *Queue) SetContentType(contentType string) {
	q.contentType = contentType
}

// Produce ...
func (q *Queue) Produce(message []byte) error {
	return q.channel.Publish(
		"",     // exchange
		q.name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  q.contentType,
			Body:         message,
		},
	)
}
