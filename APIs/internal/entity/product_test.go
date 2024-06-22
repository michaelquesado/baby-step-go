package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("any-product", 10)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "any-product", p.Name)
	assert.Equal(t, 10.0, p.Price)
}

func TestProductNameValidation(t *testing.T) {
	p, err := NewProduct("", 0)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrRequiredName, err)
}

func TestProductPriceValidation(t *testing.T) {
	p, err := NewProduct("any-prod", -1)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, ErrRequiredPrice, err)
}
