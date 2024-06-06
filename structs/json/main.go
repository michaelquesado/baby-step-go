package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Nome  string
	Saldo float32
	Pais  string `json:"pais"`
}

func main() {

	conta := Conta{Nome: "BancoX", Saldo: 150.25}
	json.NewEncoder(os.Stdout).Encode(conta)

	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(res)
	println(string(res))

	contaComoJson := []byte(`{"Nome": "BancoY", "Saldo": 12.23, "pais": "BRL" }`)
	var contaY Conta
	err = json.Unmarshal(contaComoJson, &contaY)
	if err != nil {
		panic(err)
	}
	fmt.Print(contaY, "\n")

}
