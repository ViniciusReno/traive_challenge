package controllers

import (
	"context"
	"net/http"

	"github.com/ViniciusReno/traive/internal/models"
	accountrepo "github.com/ViniciusReno/traive/internal/repository"
	"github.com/ViniciusReno/traive/internal/sqs"
	"github.com/sirupsen/logrus"
)

const queueName = "transaction-notification"

type TransactionController struct {
	db *accountrepo.Repository
}

func NewTransactionController(db *accountrepo.Repository) *TransactionController {
	return &TransactionController{
		db: db,
	}
}

func (tc TransactionController) CreateTransactions(request []models.Transaction) (int, error) {
	err := tc.db.CreateTransactions(request)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	go func() {
		sqsSenderService := sqs.NewService()
		_, err = sqsSenderService.SendMessage(context.TODO(), queueName, request)
		if err != nil {
			logrus.Error("error sending msg to sqs")
		}
	}()

	return http.StatusOK, nil
}

func (uc TransactionController) ListTransactionsByID(id string) (models.TransactionResponse, int, error) {
	transaction, err := uc.db.ListTransactionsByID(id)
	if err != nil {
		return models.TransactionResponse{}, http.StatusInternalServerError, err
	}

	return transaction, http.StatusOK, nil
}

func (tc TransactionController) ListTransactions(page int, limit int, filters map[string]interface{}) ([]models.TransactionResponse, int, error) {
	Transactions, err := tc.db.ListTransactions(page, limit, filters)
	if err != nil {
		return []models.TransactionResponse{}, http.StatusInternalServerError, err
	}

	if len(Transactions) == 0 {
		return Transactions, http.StatusNoContent, nil
	}

	return Transactions, http.StatusOK, nil
}
