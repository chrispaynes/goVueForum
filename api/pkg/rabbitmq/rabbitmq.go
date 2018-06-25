package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Init ...
func Init(url string) {
	var err error
	// Initialize the package level "conn" variable that represents the connection the the rabbitmq server
	conn, err = amqp.Dial(url)
	failOnError(err, "failed to connect to RabbitMQ")
	fmt.Println("CONNECTED TO RABBIT ", conn)
	defer conn.Close()

	if err := Publish("TEST123", []byte("HELLO WORLD")); err != nil {
		failOnError(err, "\nfailed to publish message\n")
	}
}

// Publish ...
func Publish(s string, msg []byte) error {
	// create a channel through which we publish
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         []byte("HELLO WORLD"),
		})
	failOnError(err, "Failed to publish a message")

	return nil
}
