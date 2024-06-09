package main

import (
	"net/http"
	"strings"
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
		t := template.New("template.html")
		t.Funcs(template.FuncMap{"Upper": strings.ToUpper})
		t = template.Must(t.ParseFiles("template.html"))
		err := t.Execute(w, Books{
			{Name: "harry potter", ReleaseDate: time.Now()},
			{Name: "human history", ReleaseDate: time.Now()}})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
