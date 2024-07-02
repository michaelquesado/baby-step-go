package main

import (
	"fmt"
	"time"
)

func worker(name int, data chan int) {
	for x := range data {
		fmt.Printf("worker: %d process %d \n", name, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	for i := 0; i < 100; i++ {
		go worker(i, data)
	}

	for i := 0; i < 1000; i++ {
		data <- i
	}
}
