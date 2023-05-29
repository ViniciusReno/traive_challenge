package mock

import (
	"github.com/ViniciusReno/traive/internal/models"
)

func (m *MockRepository) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockRepository) Login(req models.LoginRequest) (models.User, error) {
	args := m.Called(req)
	return args.Get(0).(models.User), args.Error(1)
}
