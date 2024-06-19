package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Studant struct {
	ID   int `gorm:"primaryKey"`
	Name string
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/gotest?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(db)
	}

	db.AutoMigrate(&Studant{})

	tx := db.Begin()
	err = tx.Debug().Create(&Studant{Name: "Carla"}).Error
	if err != nil {
		tx.Debug().Rollback()
		panic(err)
	}
	tx.Debug().Commit()

}
