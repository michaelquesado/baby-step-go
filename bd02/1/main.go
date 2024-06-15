package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	Id    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		Id:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/gotest")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
