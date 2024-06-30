package main

import (
	"net/http"
	"sync/atomic"
)

var number int64 = 0

func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// number++
		atomic.AddInt64(&number, 1)
		// m.Unlock()
		w.WriteHeader(201)
	})
	http.ListenAndServe(":8080", nil)
}
