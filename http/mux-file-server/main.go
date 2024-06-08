package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./public/")))
	http.ListenAndServe(":8080", mux)
}
