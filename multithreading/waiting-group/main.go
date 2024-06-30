package main

import (
	"fmt"
	"sync"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Printf("taks[%d] at %s\n", i, name)
		wg.Done()
	}
}
func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	go task("A", &wg)
	go task("B", &wg)
	wg.Wait()
}
