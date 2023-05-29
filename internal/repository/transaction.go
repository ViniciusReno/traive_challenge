package repository

import (
	"errors"
	"fmt"

	"github.com/ViniciusReno/traive/internal/models"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (repo *Repository) CreateTransactions(Transactions []models.Transaction) error {
	for _, Transaction := range Transactions {
		err := repo.database.Create(&Transaction).Error
		if err != nil {
			logrus.Error(err)
			return fmt.Errorf("error creating Transaction")
		}
	}

	return nil
}

func (repo *Repository) ListTransactionsByID(id string) (models.TransactionResponse, error) {
	t := models.Transaction{}

	err := repo.database.Where("id = ?", id).First(&t).Error
	if err != nil {
		logrus.Warn(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.TransactionResponse{}, fmt.Errorf("transaction not found for ID %v", id)
		}
		return models.TransactionResponse{}, fmt.Errorf("error getting transaction: %v", err)
	}

	responseTransaction := models.TransactionResponse{
		Amount:        t.Amount,
		CreatedAt:     t.CreatedAt,
		OperationType: t.OperationType,
		Origin:        t.Origin,
		TransactionID: uuid.UUID(t.ID),
		UpdatedAt:     t.UpdatedAt,
		UserID:        t.UserID,
	}

	return responseTransaction, nil
}

func (repo *Repository) ListTransactions(page int, limit int, filters map[string]interface{}) ([]models.TransactionResponse, error) {
	offset := (page - 1) * limit
	query := repo.database.Model(&models.Transaction{}).Offset(offset).Limit(limit)

	if filters != nil {
		query = query.Where(filters)
	}

	transfers := []models.Transaction{}
	if err := query.Find(&transfers).Error; err != nil {
		logrus.Error(err)
		return []models.TransactionResponse{}, fmt.Errorf("error getting transactions: %v", err)
	}

	transactionResponse := []models.TransactionResponse{}
	for _, v := range transfers {
		tr := models.TransactionResponse{
			Amount:        v.Amount,
			CreatedAt:     v.CreatedAt,
			OperationType: v.OperationType,
			Origin:        v.Origin,
			TransactionID: uuid.UUID(v.ID),
			UpdatedAt:     v.UpdatedAt,
			UserID:        v.UserID,
		}

		transactionResponse = append(transactionResponse, tr)
	}

	return transactionResponse, nil
}
