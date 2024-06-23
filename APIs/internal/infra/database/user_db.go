package database

import (
	"github.com/michaelquesado/baby-step-go/APIs/internal/entity"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *UserRepo) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
