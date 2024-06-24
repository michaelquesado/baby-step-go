package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/michaelquesado/baby-step-go/APIs/internal/dto"
	"github.com/michaelquesado/baby-step-go/APIs/internal/entity"
	"github.com/michaelquesado/baby-step-go/APIs/internal/infra/database"
	pkg "github.com/michaelquesado/baby-step-go/APIs/pkg/entity"
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
	w.WriteHeader(http.StatusCreated)
}

func (p *ProductHandler) FindOneProductHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := p.Repo.FindOne(id)

	if err != nil && err.Error() != "record not found" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err != nil && err.Error() == "record not found" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (p *ProductHandler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parsedId, err := pkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	var product entity.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product.ID = parsedId
	err = p.Repo.Update(&product)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (p *ProductHandler) ListAllProductHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	perPage := r.URL.Query().Get("per_page")
	sort := r.URL.Query().Get("sort")
	println("perPage: " + perPage)
	if page == "" {
		page = "1"
	}
	if perPage == "" {
		perPage = "10"
	}
	parsedPage, err := strconv.Atoi(page)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	parsedPerPage, err := strconv.Atoi(perPage)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	products, err := p.Repo.FindAll(parsedPage, parsedPerPage, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}