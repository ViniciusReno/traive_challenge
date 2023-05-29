package mock

import "github.com/ViniciusReno/traive/internal/models"

func (m *MockRepository) CreateTransactions(transactions []models.Transaction) error {
	args := m.Called(transactions)
	return args.Error(0)
}

func (m *MockRepository) ListTransactionsByID(id string) ([]models.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).([]models.Transaction), args.Error(1)
}

func (m *MockRepository) ListTransactions(page int, limit int, filters map[string]interface{}) ([]models.TransactionResponse, error) {
	args := m.Called(page, limit, filters)
	return args.Get(0).([]models.TransactionResponse), args.Error(1)
}
