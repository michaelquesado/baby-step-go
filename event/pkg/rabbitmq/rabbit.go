package rabbitmq

import "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp091.Channel, error) {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}

func Consume(ch *amqp091.Channel, out chan<- amqp091.Delivery) error {
	msgs, err := ch.Consume(
		"fila-brasileiro",
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	for msg := range msgs {
		out <- msg
	}
	return nil
}

func Producer(ch *amqp091.Channel, msg []byte, exchange string) error {
	err := ch.Publish(
		exchange,
		"",
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
