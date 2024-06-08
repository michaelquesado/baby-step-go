package main

import "net/http"

func main() {

	http.HandleFunc("/", helloWordHandler)
	http.ListenAndServe(":8080", nil)
}

func helloWordHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !r.URL.Query().Has("cep") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte("Hello, Word!"))
}
