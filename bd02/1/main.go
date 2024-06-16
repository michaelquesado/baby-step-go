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
	println("All products 👇")
	products, err := findAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, pro := range products {
		fmt.Printf("id: %v | name: %v | price: %2f\n", pro.Id, pro.Name, pro.Price)
	}
	if err := deleteProduct(db, p.Id); err != nil {
		panic(err)
	}
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

func findAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select * from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.Id, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
