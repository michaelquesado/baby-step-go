package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 100         //write value to the channel
	fmt.Println(<-ch) //read value from the channel
}
