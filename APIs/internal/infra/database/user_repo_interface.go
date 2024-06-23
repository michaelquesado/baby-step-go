package database

import "github.com/michaelquesado/baby-step-go/APIs/internal/entity"

type UserRepoInterface interface {
	Create(u *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
