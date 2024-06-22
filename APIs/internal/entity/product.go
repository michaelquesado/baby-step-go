package entity

import (
	"errors"
	"time"

	"github.com/michaelquesado/baby-step-go/APIs/pkg/entity"
)

var (
	ErrRequiredID    error = errors.New("ID is required")
	ErrInvalidID     error = errors.New("invalid ID")
	ErrRequiredName  error = errors.New("name is required")
	ErrRequiredPrice error = errors.New("price is required")
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
		return ErrRequiredID
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}
	if p.Name == "" {
		return ErrRequiredName
	}
	if p.Price <= 0 {
		return ErrRequiredPrice
	}
	return nil
}
