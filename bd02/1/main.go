package main

import (
	"database/sql"
	"fmt"

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
	p := NewProduct("Camisa Polo", 159.29)
	if err := insertProduct(db, p); err != nil {
		panic(err)
	}
	p.Name = "Camisa Super Polo 2 "
	p.Price = 99.99
	if err := updateProduct(db, p); err != nil {
		panic(err)
	}

	product, err := findOneProduct(db, p.Id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product name %v and price %2f \n", product.Name, product.Price)
}

func insertProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("insert into products (id, name, price) values (?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.Id, p.Name, p.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.Name, p.Price, p.Id)
	if err != nil {
		return err
	}
	return nil
}

func findOneProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p Product
	err = stmt.QueryRow(id).Scan(&p.Id, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
