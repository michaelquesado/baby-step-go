package main

import (
	"net/http"
	"text/template"
	"time"
)

type Book struct {
	Name        string
	ReleaseDate time.Time
}

type Books []Book

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Books{
			{Name: "Harry Potter", ReleaseDate: time.Now()},
			{Name: "Human History", ReleaseDate: time.Now()}})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
