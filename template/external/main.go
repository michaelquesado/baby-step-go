package main

import (
	"os"
	"text/template"
)

type Game struct {
	Name string
	Type string
}

type Games []Game

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Games{{Name: "The Last Of Us I", Type: "Survive"}, {Name: "A link to the past", Type: "Adventure"}})
	if err != nil {
		panic(err)
	}
}
