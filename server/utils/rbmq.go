package utils

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var rbmqconnect *amqp.Connection
var rbmqchann *amqp.Channel

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func InitRmq() {
	// 建立连接
	rbmqconnect, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	failOnError(err, "Failed to connect to RabbitMQ")
	// 打开频道
	rbmqchann, err := rbmqconnect.Channel()
	failOnError(err, "Failed to open a channel")
	err = rbmqchann.ExchangeDeclare(
		"order.dispatch", // name
		"fanout",         // type
		false,            // durability
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	failOnError(err, "Failed to declare a exchange")
	defer rbmqchann.Close()
}

func Send(data []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := rbmqchann.PublishWithContext(ctx,
		"order.dispatch", // exchange
		"",               // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("12"),
		})
	if err != nil {
		return err
	}
	return nil
}

func Recevie() {

}

func CloseConnect() {
	defer rbmqconnect.Close()
}
