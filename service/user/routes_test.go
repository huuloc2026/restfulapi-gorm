package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/huuloc2026/restfulapi-gorm.git/types"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should be fail if the user payload invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			Name:     "user",
			Email:    "client01",
			Password: "password123",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		request := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(request, req)
		if request.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d,got %d", http.StatusBadRequest, request.Code)
		}
		expectedResponse := `{"error":"invalid payload"}`
		if request.Body.String() != expectedResponse {
			t.Errorf("expected response %s, got %s", expectedResponse, request.Body.String())
		}
	})
	t.Run("should correctly register user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			Name:     "user",
			Email:    "client01@gmail.com",
			Password: "password123",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		request := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(request, req)
		if request.Code != http.StatusCreated {
			t.Errorf("expected status code %d,got %d", http.StatusCreated, request.Code)
		}
	})

}

type mockUserStore struct {
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("User not found")
}
