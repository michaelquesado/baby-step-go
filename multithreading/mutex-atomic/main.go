package main

import (
	"net/http"
	"sync"
)

var number int64 = 0

func main() {
	m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		number++
		m.Unlock()
		w.WriteHeader(201)
	})
	http.ListenAndServe(":8080", nil)
}
