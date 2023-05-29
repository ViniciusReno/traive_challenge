package repository

import (
	"reflect"
	"testing"

	"github.com/ViniciusReno/traive/internal/models"
	"github.com/ViniciusReno/traive/internal/repository/mock"
)

func TestCreateUser(t *testing.T) {
	repo := &mock.MockRepository{}
	user := &models.User{}

	repo.On("CreateUser", user).Return(nil)

	err := repo.CreateUser(user)

	repo.AssertCalled(t, "CreateUser", user)

	if err != nil {
		t.Errorf("Expected error: nil, Received error: %v", err)
	}
}

func TestLogin(t *testing.T) {
	repo := &mock.MockRepository{}
	req := models.LoginRequest{}
	expectedUser := models.User{}

	repo.On("Login", req).Return(expectedUser, nil)

	user, err := repo.Login(req)

	repo.AssertCalled(t, "Login", req)

	if err != nil {
		t.Errorf("Expected error: nil, Received error: %v", err)
	}

	if !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("Expected user: %+v, Received user: %+v", expectedUser, user)
	}
}
