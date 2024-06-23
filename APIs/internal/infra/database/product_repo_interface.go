package database

import "github.com/michaelquesado/baby-step-go/APIs/internal/entity"

type ProductRepoInterface interface {
	Create(product *entity.Product) error
	FindOne(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
	FindAll(page, per_page int, sort string) ([]entity.Product, error)
}
