package main

import (
	"app/config"
	"app/internal/jobs"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	amqp "github.com/rabbitmq/amqp091-go"
)

var jobRegistry = map[string]func([]byte) error{
	"Test":       jobs.Test,
	"HelloWorld": jobs.HelloWorld,
}

var conn *amqp.Connection
var channel *amqp.Channel
var f *os.File

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func initLogging() {
	logDir := "log"
	logFile := "worker.log"
	logPath := filepath.Join(logDir, logFile)

	file, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	iw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(iw)
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

	conn, err = amqp.Dial(dsn)
	failOnError(err, "Failed to connect to RabbitMQ")

	channel, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
}

func main() {
	initLogging()
	defer f.Close()

	initRabbitMQ()
	defer conn.Close()
	defer channel.Close()

	err := channel.ExchangeDeclare(
		"jobs",   // exchange
		"direct", // type
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := channel.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := channel.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			fmt.Printf("🟢 Received jobs [%s]: %s\n", d.RoutingKey, d.Body)

			if handlers, ok := jobRegistry[d.RoutingKey]; ok {
				err := handlers(d.Body)
				if err != nil {
					log.Printf("jobs %s failed: %v\n", d.RoutingKey, err)
					d.Nack(false, false) // Reject the message and do not requeue
				}
				log.Printf("Received a message: %s", d.Body)
				log.Printf("Done")
				d.Ack(false)
			} else {
				log.Println("Unknown jobs action:", d.RoutingKey)
				d.Nack(false, false) // Reject the message and do not requeue
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	select {}
}
