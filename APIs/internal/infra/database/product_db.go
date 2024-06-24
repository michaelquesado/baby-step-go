package database

import (
	"github.com/michaelquesado/baby-step-go/APIs/internal/entity"
	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{DB: db}
}

func (p *ProductRepo) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *ProductRepo) FindOne(id string) (*entity.Product, error) {
	var product entity.Product
	if err := p.DB.First(&product, " id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil

}

func (p *ProductRepo) FindAll(page, perPage int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && perPage != 0 {
		err = p.DB.Limit(perPage).Offset((page - 1) * perPage).Order("created_at " + sort).Find(&products).Error
		return products, err
	}
	err = p.DB.Order("created_at " + sort).Find(&products).Error
	return products, err
}

func (p *ProductRepo) Update(product *entity.Product) error {
	if _, err := p.FindOne(product.ID.String()); err != nil {
		return err
	}
	return p.DB.Updates(&product).Error
}

func (p *ProductRepo) Delete(id string) error {
	product, err := p.FindOne(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}
