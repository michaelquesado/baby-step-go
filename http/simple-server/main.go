package main

import "net/http"

func main() {

	http.HandleFunc("/", helloWord)
	http.ListenAndServe(":8080", nil)
}

func helloWord(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Word!"))
}
