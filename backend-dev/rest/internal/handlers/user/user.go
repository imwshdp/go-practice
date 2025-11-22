package user

import (
	"fmt"
	"net/http"
	"rest/internal/dto"
	"rest/internal/models"
	user "rest/internal/storage/postgres/user"
	httpUtils "rest/internal/utils/http"
	validateUtil "rest/internal/utils/validator"

	"rest/internal/services/auth"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type userHandler struct {
	userRepo user.UserRepository
}

func NewUserHandler(userRepo user.UserRepository) *userHandler {
	return &userHandler{
		userRepo: userRepo,
	}
}

func (h *userHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.login).Methods("POST")
	router.HandleFunc("/register", h.register).Methods("POST")
}

func (h *userHandler) login(w http.ResponseWriter, r *http.Request) {

}

func (h *userHandler) register(w http.ResponseWriter, r *http.Request) {
	var payload models.RegisterUserPayload

	if err := httpUtils.ParseJSON(r, &payload); err != nil {
		httpUtils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validation
	if err := validateUtil.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		httpUtils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid payload %v", errors))
		return
	}

	_, err := h.userRepo.GetByEmail(payload.Email)
	if err == nil {
		httpUtils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		httpUtils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.userRepo.Create(
		dto.User{
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
			Email:     payload.Email,
			Password:  hashedPassword,
		},
	)
	if err != nil {
		httpUtils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	httpUtils.WriteJSON(w, http.StatusCreated, nil)
}
