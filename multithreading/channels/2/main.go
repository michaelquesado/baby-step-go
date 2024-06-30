package main

func main() {

	cn := make(chan int)
	go publisher(cn)
	for x := range cn {
		println(x)
	}
}

func publisher(cn chan int) {
	for i := 0; i < 10; i++ {
		cn <- i
	}
	close(cn)
}
