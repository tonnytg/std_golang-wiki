package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

const (
	amqpServer = "amqps://wwwwwwwww:xxxxxxx@yyyyyyy.rmq.cloudamqp.com/zzzzzzzzz"
)

func sendMessageToAMQP(msg string) error {

	conn, err := amqp.Dial(amqpServer)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
	}

	q, err := ch.QueueDeclare(
		"test", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}

	body := msg
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to publish a message", err)
	}

	log.Printf("Sent successfully")

	return nil
}

func getMessageFromAMQP() (string, error) {

	conn, err := amqp.Dial(amqpServer)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
	}

	q, err := ch.QueueDeclare(
		"test", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to register a consumer", err)
	}

	var message string
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		message = string(d.Body)
		break
	}

	conn.Close()
	return message, nil
}

func main() {
	fmt.Println("Start send message to AMQP")
	msgToSend := `{"email":"test@test.com", "status": "active"}`
	err := sendMessageToAMQP(msgToSend)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("start get same message from AMQP")
	msgReceived, err := getMessageFromAMQP()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if msgToSend == msgReceived {
		fmt.Println("Success: message received is same as sent")
	} else {
		fmt.Println("Error: message received is not same as sent")
	}
}
