package queue

import (
	"time"

	"github.com/mjah/jwt-auth/logger"
)

type messageConsumer func([]byte)

// Consume ...
func (q *Queue) Consume(consumer messageConsumer) {
	for {
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
			time.Sleep(5 * time.Second)
			logger.Log().Info("Attempting consumer recovery.")
			continue
		}
		for d := range msgs {
			consumer(d.Body)
			d.Ack(false)
		}
	}
}
