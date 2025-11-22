package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"rest/internal/dto"
	"rest/internal/models"
	"testing"

	"github.com/gorilla/mux"
)

type mockUserRepo struct{}

func (r *mockUserRepo) GetByEmail(email string) (*dto.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (r *mockUserRepo) GetById(id int) (*dto.User, error) {
	return nil, nil
}

func (r *mockUserRepo) Create(user dto.User) error {
	return nil
}

func TestUserServiceHandlers(t *testing.T) {
	userRepo := &mockUserRepo{}
	handler := NewUserHandler(userRepo)

	t.Run("register user payload is invalid", func(t *testing.T) {
		payload, _ := json.Marshal(models.RegisterUserPayload{
			FirstName: "User",
			LastName:  "Smith",
			Email:     "invalid",
			Password:  "123456",
		})

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/register", handler.register)
		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, recorder.Code)
		}
	})

	t.Run("register user payload is valid", func(t *testing.T) {
		payload, _ := json.Marshal(models.RegisterUserPayload{
			FirstName: "User",
			LastName:  "Smith",
			Email:     "test_user@gmail.com",
			Password:  "123456",
		})

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/register", handler.register)
		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, recorder.Code)
		}
	})
}
