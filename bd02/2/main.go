package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Company struct {
	Id   int `gorm:"primaryKey;autoIncrement:true"`
	Name string
	CNPJ string
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/gotest"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Company{})

	// db.Create(&Company{
	// 	Name: "UPA",
	// 	CNPJ: "87278810001/02",
	// })

	// companies := []Company{
	// 	{Name: "FNJ", CNPJ: "0019283820001/02"},
	// 	{Name: "Paraiso", CNPJ: "991923902010001/93"},
	// }
	// db.Create(&companies)

	// var c Company

	// db.First(&c, 1)
	// fmt.Println(c)

	// db.First(&c, "name = ?", "Paraiso")
	// fmt.Println(c)

	// var companies []Company
	// db.Find(&companies)
	// for _, c := range companies {
	// 	fmt.Println(c)
	// }

	var companies []Company
	db.Where("name like ?", "%F%").Find(&companies)
	for _, c := range companies {
		fmt.Println(c)
	}
}
