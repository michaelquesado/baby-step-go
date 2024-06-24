package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/michaelquesado/baby-step-go/APIs/internal/dto"
	"github.com/michaelquesado/baby-step-go/APIs/internal/entity"
	"github.com/michaelquesado/baby-step-go/APIs/internal/infra/database"
)

type UserHandler struct {
	UserRepo database.UserRepoInterface
}

func NewUserHandler(repo database.UserRepoInterface) *UserHandler {
	return &UserHandler{UserRepo: repo}
}

func (uh *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var dto dto.UserDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	user, err := entity.NewUser(dto.Name, dto.Email, dto.Password)
	log.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = uh.UserRepo.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
