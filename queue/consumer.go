package queue

type messageConsumer func([]byte)

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
			consumer(d.Body)
			d.Ack(false)
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

// Consume ...
func (q *Queue) Consume(consumer messageConsumer) error {
	q.consumers = append(q.consumers, consumer)

	if err := q.registerConsumer(consumer); err != nil {
		return err
	}

	return nil
}
