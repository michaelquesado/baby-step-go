package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/michaelquesado/baby-step-go/APIs/configs"
	"github.com/michaelquesado/baby-step-go/APIs/internal/entity"
	"github.com/michaelquesado/baby-step-go/APIs/internal/infra/database"
	"github.com/michaelquesado/baby-step-go/APIs/internal/infra/webserver/handlers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config, _ := configs.LoadConfig(".")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productRepo := database.NewProductRepo(db)
	handler := handlers.NewProductHandler(productRepo)

	userRepo := database.NewUserRepo(db)
	userHandler := handlers.NewUserHandler(userRepo)

	r := chi.NewRouter()
	r.Post("/product", handler.CreateProductHandler)
	r.Get("/product/{id}", handler.FindOneProductHandler)
	r.Patch("/product/{id}", handler.UpdateProductHandler)
	r.Get("/product", handler.ListAllProductHandler)

	r.Post("/user", userHandler.CreateUserHandler)

	http.ListenAndServe(config.WebServerPort, r)
}
