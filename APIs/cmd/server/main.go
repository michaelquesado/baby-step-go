package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
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

	userJwtHandler := handlers.NewUserJwtHandler(*userRepo, config.TokenAuth, config.JWTExperesIn)

	r := chi.NewRouter()
	r.Route("/product", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", handler.CreateProductHandler)
		r.Get("/{id}", handler.FindOneProductHandler)
		r.Patch("/{id}", handler.UpdateProductHandler)
		r.Get("/", handler.ListAllProductHandler)
	})

	r.Post("/user", userHandler.CreateUserHandler)
	r.Post("/user/login", userJwtHandler.GenerateTokenHandler)

	http.ListenAndServe(config.WebServerPort, r)
}
