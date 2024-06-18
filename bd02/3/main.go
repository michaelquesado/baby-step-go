package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID      int `gorm:"primaryKey"`
	Name    string
	Product []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/gotest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// c := Category{Name: "Brinquedo"}
	// db.Create(&c)

	// p := Product{Name: "Caminhao", Price: 999, CategoryID: c.ID}
	// db.Create(&p)

	// sN := SerialNumber{Number: "9912831", ProductID: p.ID}
	// db.Create(&sN)

	// var products []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products)

	// for _, product := range products {
	// 	fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	// }

	var categories []Category
	err = db.Model(&Category{}).Preload("Product.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Product {
			fmt.Println("-", product.Name, product.SerialNumber.Number)
		}
	}
}
