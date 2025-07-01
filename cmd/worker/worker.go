package main

import (
	"app/config"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Conn *amqp.Connection
var Channel *amqp.Channel

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

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

	Conn, err = amqp.Dial(dsn)
	failOnError(err, "Failed to connect to RabbitMQ")

	Channel, err = Conn.Channel()
	failOnError(err, "Failed to open a channel")
}

func main() {
	initRabbitMQ()
	defer Conn.Close()
	defer Channel.Close()

	q, err := Channel.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = Channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := Channel.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			log.Printf("Done")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
