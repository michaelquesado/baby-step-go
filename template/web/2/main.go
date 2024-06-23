package main

import (
	"os"
	"text/template"
)

type Phone struct {
	Name string
	Os   string
}

type Phones []Phone

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(os.Stdout, Phones{{Name: "Iphone 11", Os: "IOS"}, {Name: "MotoG", Os: "Android"}})
	if err != nil {
		panic(err)
	}
}
