package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()
	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- 1
	}()

	select {
	case msg := <-ch1:
		fmt.Printf("channel 1: %d \n", msg)
	case msg := <-ch2:
		fmt.Printf("channel 2: %d \n", msg)
	}
}
