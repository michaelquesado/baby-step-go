package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
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

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		var data Cep
		err = json.Unmarshal(res, &data)
		if err != nil {
			panic(err)
		}
		println(data.Logradouro)
		file, err := os.Create("cep.json")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		decode, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		file.WriteString(string(decode))
	}
}
