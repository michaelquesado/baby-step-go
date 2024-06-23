package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/michaelquesado/baby-step-go/APIs/configs"
	"github.com/michaelquesado/baby-step-go/APIs/internal/dto"
	"github.com/michaelquesado/baby-step-go/APIs/internal/entity"
	"github.com/michaelquesado/baby-step-go/APIs/internal/infra/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config, _ := configs.LoadConfig(".")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productRepo := database.NewProductRepo(db)
	handler := NewProductHandler(productRepo)

	http.HandleFunc("POST /product", handler.CreateProductHandler)
	http.ListenAndServe(config.WebServerPort, nil)
}

type ProductHandler struct {
	Repo database.ProductRepoInterface
}

func NewProductHandler(repo database.ProductRepoInterface) *ProductHandler {
	return &ProductHandler{Repo: repo}
}

func (p *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var dto dto.ProductDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := entity.NewProduct(dto.Name, dto.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = p.Repo.Create(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
