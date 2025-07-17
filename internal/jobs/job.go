package jobs

import (
	"app/config"
	"app/util"
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var conn *amqp.Connection
var channel *amqp.Channel

func initRabbitMQ() {
	dsn := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/%s",
		config.Config("RABBITMQ_USER"),
		config.Config("RABBITMQ_PASSWORD"),
		config.Config("RABBITMQ_HOST"),
		config.Config("RABBITMQ_PORT"),
		config.Config("RABBITMQ_VHOST"),
	)
	var err error

	conn, err = amqp.Dial(dsn)
	util.FailOnError(err, "Failed to connect to RabbitMQ")

	channel, err = conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
}

func PublishJob(action string, body string) {
	initRabbitMQ()
	defer conn.Close()
	defer channel.Close()

	err := channel.ExchangeDeclare(
		"jobs",   // exchange
		"direct", // type
		true,     // durable
		false,
		false,
		false,
		nil,
	)
	util.FailOnError(err, "Failed to declare a queue")
	q, err := channel.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	util.FailOnError(err, "Failed to declare a queue")

	log.Printf("Declared queue: %s", q.Name)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = channel.PublishWithContext(
		ctx,
		"jobs", // exchange
		action, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		},
	)

	util.FailOnError(err, "Failed to publish a message")
}
