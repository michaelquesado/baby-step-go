package main

import (
	"michaelquesado/baby-step-go/event/pkg/rabbitmq"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	msgs := make(chan amqp091.Delivery)
	go rabbitmq.Consume(ch, msgs)
	for msg := range msgs {
		println(string(msg.Body))
		msg.Ack(false)
	}
}
