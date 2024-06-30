package main

import "fmt"

func main() {
	channel := make(chan string)

	go func() {
		channel <- "any message"
	}()

	msg := <-channel

	fmt.Println(msg)
}
