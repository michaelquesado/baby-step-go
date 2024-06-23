package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/michaelquesado/baby-step-go/APIs/internal/dto"
	"github.com/michaelquesado/baby-step-go/APIs/internal/entity"
	"github.com/michaelquesado/baby-step-go/APIs/internal/infra/database"
)

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
