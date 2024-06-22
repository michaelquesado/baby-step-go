package entity

import (
	"github.com/michaelquesado/baby-step-go/APIs/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, pass string) (*User, error) {
	hash, err := encryptPass(pass)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: hash,
	}, nil
}

func encryptPass(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (u *User) ValidadePassword(pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
	return err == nil
}
