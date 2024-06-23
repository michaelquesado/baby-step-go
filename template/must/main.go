package main

import (
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Quesado", Age: 29}
	err := template.Must(template.New("PersonTemplate").Parse("O cliente {{.Name}} tem {{.Age}} anos\n")).Execute(os.Stdout, p)
	if err != nil {
		panic(err)
	}

}
