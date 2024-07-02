package main

import (
	"fmt"
	"slices"
)

func main() {
	// var foo string = "foo"
	// bar := "bar"

	// fmt.Println(foo, bar)

	// var (
	// 	foobar string = "foobar"
	// 	barfoo string = "barfoo"
	// )
	// fmt.Println(foobar, barfoo)

	// another, some := "another", "some"
	// fmt.Println(another, some)

	// //chage value
	// another, some = foobar, barfoo
	// fmt.Println(another, some)

	arr := []int{1, 15, 35, 2, 11}
	slices.Sort(arr)

	fmt.Println(arr)

}
