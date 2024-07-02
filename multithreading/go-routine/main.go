package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d task %s\n", i, name)
		time.Sleep(1 * time.Second)
	}
}
func main() {
	go task("A")
	go task("B")
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d task %s\n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(10 * time.Second)
}
