package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	http.HandleFunc("/", helloWordHandler)
	http.ListenAndServe(":8080", nil)
}

type Validation struct {
	Error string `json:"error"`
}

func NewValidation(error string) *Validation {
	v := Validation{Error: error}
	return &v
}

type CepResponse struct {
	Msg string `json:"msg"`
}

func NewCepResponse(msg string) *CepResponse {
	c := CepResponse{Msg: msg}
	return &c
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
	w.Header().Set("Content-Type", "application/json")
	cep := r.URL.Query().Get("cep")
	if len(cep) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(string(parseStrucToJson(NewValidation("Invalid cep")))))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(parseStrucToJson(NewCepResponse("Hello, World!")))))
}

func parseStrucToJson(s any) []byte {
	res, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return res
}
