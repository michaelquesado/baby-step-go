package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	c := Curso{Nome: "Golang", CargaHoraria: 1}
	tmp := template.New("simpleTemplate")
	tmp, _ = tmp.Parse("O curos {{.Nome}} tem um total de {{.CargaHoraria}} hora(s)\n")

	err := tmp.Execute(os.Stdout, c)
	if err != nil {
		panic(err)
	}
}
