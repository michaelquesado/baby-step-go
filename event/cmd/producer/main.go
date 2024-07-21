package main

import "michaelquesado/baby-step-go/event/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	err = rabbitmq.Producer(ch, []byte("Olá mundo"), "amq.direct")
	if err != nil {
		panic(err)
	}
}
