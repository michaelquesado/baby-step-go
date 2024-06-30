package main

import "sync"

func main() {
	cn := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)
	go publisher(cn)
	go reader(cn, &wg)
	wg.Wait()
}

func reader(cn chan int, wg *sync.WaitGroup) {
	for x := range cn {
		println(x)
		wg.Done()
	}
}

func publisher(cn chan int) {
	for i := 0; i < 10; i++ {
		cn <- i
	}
	close(cn)
}
