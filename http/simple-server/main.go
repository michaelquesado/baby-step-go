package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Cep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
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
	w.Header().Set("Content-Type", "application/json")
	cep := r.URL.Query().Get("cep")
	if len(cep) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(string(parseStrucToJson(NewValidation("Invalid cep")))))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(parseStrucToJson(fetchCep(cep)))))
}

func parseStrucToJson(s any) []byte {
	res, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return res
}

func fetchCep(cep string) *Cep {
	res, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var c = Cep{}
	err = json.Unmarshal(body, &c)
	if err != nil {
		panic(err)
	}
	return &c
}
