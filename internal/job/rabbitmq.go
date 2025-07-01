package job

import (
	"app/config"
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Conn *amqp.Connection
var Channel *amqp.Channel

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func InitRabbitMQ() {

	dsn := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/%s",
		config.Config("RABBITMQ_USER"),
		config.Config("RABBITMQ_PASSWORD"),
		config.Config("RABBITMQ_HOST"),
		config.Config("RABBITMQ_PORT"),
		config.Config("RABBITMQ_VHOST"),
	)

	var err error

	Conn, err = amqp.Dial(dsn)
	failOnError(err, "Failed to connect to RabbitMQ")

	Channel, err = Conn.Channel()
	failOnError(err, "Failed to open a channel")
}

func PublishJob(queueName string, message string) {
	q, err := Channel.QueueDeclare(queueName, true, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = Channel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(message),
		},
	)
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", message)
}
