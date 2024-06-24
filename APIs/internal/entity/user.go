package entity

import (
	"errors"

	"github.com/michaelquesado/baby-step-go/APIs/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrRequiredEmail        error = errors.New("email is required")
	ErrRequiredUserName     error = errors.New("name is required")
	ErrRequiredUserPassword error = errors.New("password is required")
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
	user := User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: hash,
	}
	err = user.Validation()
	if err != nil {
		return nil, err
	}
	return &user, nil
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

func (u *User) Validation() error {
	if u.Email == "" {
		return ErrRequiredEmail
	}
	if u.Name == "" {
		return ErrRequiredUserName
	}
	if u.Password == "" {
		return ErrRequiredUserPassword
	}
	return nil
}
