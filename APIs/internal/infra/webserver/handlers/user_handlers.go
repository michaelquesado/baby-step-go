package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
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

type UserJwtHandler struct {
	userRepo    database.UserRepo
	jwtauth     *jwtauth.JWTAuth
	jwtExpireIn int
}

func NewUserJwtHandler(repo database.UserRepo, jwt *jwtauth.JWTAuth, expireIn int) *UserJwtHandler {
	return &UserJwtHandler{
		userRepo:    repo,
		jwtauth:     jwt,
		jwtExpireIn: expireIn,
	}
}

func (ujh *UserJwtHandler) GenerateTokenHandler(w http.ResponseWriter, r *http.Request) {
	var userLogin dto.UserLoginDTO
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	user, err := ujh.userRepo.FindByEmail(userLogin.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	log.Println(userLogin)
	if !user.ValidadePassword(userLogin.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, token, err := ujh.jwtauth.Encode(map[string]interface{}{
		"sub": user.ID,
		"exp": time.Now().Add(time.Second * time.Duration(ujh.jwtExpireIn)).Unix(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jwt := struct {
		AccessToken string `json:"access_token"`
	}{AccessToken: token}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jwt)

}
