package main

import "fmt"

type person struct {
	name string
	age  int8
}

type student struct {
	person
	school string
}

type Race struct {
	name  string
	color string
}
type Dog struct {
	name  string
	noise string
	race  Race
}

func main() {
	guy := student{person{"Jhon", 21}, "EEP"}
	fmt.Println(guy)

	srd := Dog{name: "Scooby", noise: "bark", race: Race{name: "danchround", color: "black"}}
	fmt.Println(srd)
}
