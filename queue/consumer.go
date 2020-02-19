package queue

import "github.com/mjah/jwt-auth/logger"

type messageConsumer func([]byte) error

func (q *Queue) registerConsumer(consumer messageConsumer) error {
	msgs, err := q.channel.Consume(
		q.name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			err := consumer(d.Body)
			if err != nil {
				logger.Log().Error(err)
			} else {
				d.Ack(false)
			}
		}
	}()

	return nil
}

func (q *Queue) recoverConsumer() error {
	for _, consumer := range q.consumers {
		if err := q.registerConsumer(consumer); err != nil {
			return err
		}
	}
	return nil
}

// Consume registers the consumer with the message-broker.
func (q *Queue) Consume(consumer messageConsumer) error {
	q.consumers = append(q.consumers, consumer)
	return q.registerConsumer(consumer)
}
