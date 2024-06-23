package main

import (
	"fmt"

	"github.com/michaelquesado/baby-step-go/packages02/1/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
	fmt.Println("Hello, world!")
}
