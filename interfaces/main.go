package main

import "fmt"

func main() {
	var minhaVar interface{} = "Michael Quesado"
	println(minhaVar.(string))
	cast, ok := minhaVar.(int)
	fmt.Printf("cast %v, ok %v\n", cast, ok)
	vaiDarError := minhaVar.(int)
	println(vaiDarError)
}
