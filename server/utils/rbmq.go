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
	var err error
	rbmqconnect, err = amqp.Dial("amqp://guest:guest@localhost:5672/")

	failOnError(err, "Failed to connect to RabbitMQ")
	// 打开频道
	rbmqchann, err = rbmqconnect.Channel()
	failOnError(err, "Failed to open a channel")
	err = rbmqchann.ExchangeDeclare(
		"order.dispatch", // name
		"fanout",         // type
		true,             // durability
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	failOnError(err, "Failed to declare a exchange")
}

func Send(exchange string, orderIdBytes []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := rbmqchann.PublishWithContext(ctx,
		exchange, // exchange
		"",       // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        orderIdBytes,
		})
	if err != nil {
		return err
	}
	return nil
}

func Recevie() {

}

func CloseConnect() {
	if rbmqchann != nil {
		rbmqchann.Close()
	}
	if rbmqconnect != nil {
		rbmqconnect.Close()
	}
}
