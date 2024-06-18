package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Company struct {
	Id   int `gorm:"primaryKey;autoIncrement:true"`
	Name string
	CNPJ string
	gorm.Model
}

type Product struct {
	Id    string `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/gotest?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Company{})

	// productName := "Camisa Polo"
	// var p Product
	// db.Find(&p, "name = ?", productName)
	// fmt.Println(p)

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

	// var companies []Company
	// db.Where("name like ?", "%F%").Find(&companies)
	// for _, c := range companies {
	// 	fmt.Println(c)
	// }

	// var c Company
	// db.First(&c, 1)
	// c.Name = "Univerdade Patativa do Assare"
	// db.Save(&c)

	// var c2 Company
	// db.First(&c2, 2)
	// fmt.Println(c2)

	// db.Delete(&c2)

}
