package event

import (
	"context"
	"fmt"
	"log"

	"github.com/apascualco/apascualco-auth/kit/event"
	"github.com/streadway/amqp"
)

// RabbitEventBus an in-memory implementation of the event.Bus.
type RabbitEventBus struct {
	user     string
	password string
	port     string
	vhost    string
}

func NewRabbitEventBus(user, password, port, vhost string) RabbitEventBus {
	return RabbitEventBus{
		user:     user,
		password: password,
		port:     port,
		vhost:    vhost,
	}
}

// Publish implements the event.Bus interface.
func (b RabbitEventBus) Publish(ctx context.Context, events []event.Event) error {
	con := connectRabbitMQ(b.user, b.password, b.port, b.vhost)
	ch := openChannel(con)
	for _, evt := range events {
		publish(evt, ch)
	}
	return nil
}

func connectRabbitMQ(user, password, port, vhost string) *amqp.Connection {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@rabbitmq:%s/%s", user, password, port, vhost))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	return conn
}

func openChannel(conn *amqp.Connection) *amqp.Channel {
	c, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	return c
}

func publish(e event.Event, channel *amqp.Channel) {
	log.Printf("Publish event echange %s type %s event %v", "auth", e.Type(), e)
	evt, _ := e.Marshaller()
	err := channel.Publish(
		"auth",
		string(e.Type()),
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         evt,
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	}
}
