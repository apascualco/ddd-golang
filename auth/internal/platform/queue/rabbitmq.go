package queue

import (
	"fmt"
	"log"

	auth "github.com/apascualco/apascualco-auth/internal"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	url   string
	queue string
}

func NewRabbitMQ(u, p, port, v, q string) RabbitMQ {
	return RabbitMQ{
		url:   fmt.Sprintf("amqp://%s:%s@rabbitmq:%s/%s", u, p, port, v),
		queue: q,
	}
}

func (r RabbitMQ) ReadMessages(e auth.EventHandler) error {
	log.Printf("Connecting to url %s and queue %s", r.url, r.queue)
	c, err := amqp.Dial(r.url)
	if err != nil {
		return err
	}
	defer c.Close()

	channel, err := c.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	delivery, err := channel.Consume(r.queue, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	f := make(chan bool)

	go func() {
		for d := range delivery {
			if err := e.EventHandler(string(d.Body)); err != nil {
				// Handler error
				log.Printf("Error handling event %s", err.Error())
			} else {
				d.Ack(false)
			}

		}
	}()
	<-f
	return nil
}
