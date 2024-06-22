package entity

import (
	"errors"
	"time"

	"github.com/michaelquesado/baby-step-go/APIs/pkg/entity"
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	p := Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	if err := p.Validate(); err != nil {
		return nil, err
	}
	return &p, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return errors.New("ID is required")
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return errors.New("invalid ID")
	}
	if p.Name == "" {
		return errors.New("name is required")
	}
	if p.Price <= 0 {
		return errors.New("price is required")
	}
	return nil
}
