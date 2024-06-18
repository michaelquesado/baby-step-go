package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/gotest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// c := Category{Name: "Eletronico"}
	// db.Create(&c)

	// p := Product{Name: "Geladeira", Price: 999, CategoryID: c.ID}
	// db.Create(&p)

	var products []Product
	db.Preload("Category").Find(&products)

	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)
	}
}
