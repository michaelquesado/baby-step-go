package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/michaelquesado/baby-step-go/APIs/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("::filename:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	p, _ := entity.NewProduct("any-product", 99.9)

	repo := NewProductRepo(db)

	err = repo.Create(p)

	assert.Nil(t, err)

	productExists, _ := repo.FindOne(p.ID.String())

	assert.Equal(t, productExists.ID, p.ID)
}

func TestFindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("::filename:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	repo := NewProductRepo(db)
	for i := 1; i < 24; i++ {
		product, _ := entity.NewProduct(fmt.Sprintf("any-produc: %v", i), rand.Float64()*100)
		fmt.Println(product)
		repo.DB.Create(product)
	}

	products, err := repo.FindAll(1, 10, "asc")

	fmt.Println(len(products))
	assert.Nil(t, err)
	assert.Equal(t, len(products), 10)
}

func TestUpdate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("::filename:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	p, _ := entity.NewProduct("any-product x", 99.9)

	repo := NewProductRepo(db)

	_ = repo.Create(p)

	p.Name = "any-product y"

	err = repo.Update(p)

	updatedProduct, _ := repo.FindOne(p.ID.String())

	assert.Nil(t, err)
	assert.Equal(t, p.Name, updatedProduct.Name)
}

func TestDelete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("::filename:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	p, _ := entity.NewProduct("any-product x", 99.9)

	repo := NewProductRepo(db)

	_ = repo.Create(p)

	err = repo.Delete(p.ID.String())

	assert.Nil(t, err)

	_, err = repo.FindOne(p.ID.String())

	assert.NotNil(t, err)

}
