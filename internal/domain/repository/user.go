package repository

import (
	"hospital-management-system/internal/domain/models"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByID(id int) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	Update(user *models.User) error
	Delete(id int) error
	FindAll() ([]models.User, error)
}
