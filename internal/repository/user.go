package repository

import (
	"fmt"

	"github.com/ViniciusReno/traive/internal/models"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) CreateUser(user *models.User) error {
	err := repo.database.Create(&user).Error
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("error creating user")
	}

	return nil
}

func (repo *Repository) Login(req models.LoginRequest) (models.User, error) {
	var user models.User
	if err := repo.database.Where("username = ?", req.Username).First(&user).Error; err != nil {
		logrus.Error(err)
		return models.User{}, fmt.Errorf("internal error")
	}

	return user, nil
}
