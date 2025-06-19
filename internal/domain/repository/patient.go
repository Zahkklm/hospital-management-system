package repository

import (
	"hospital-management-system/internal/domain/models"
)

// PatientRepository defines the methods for interacting with patient data.
type PatientRepository interface {
	Create(patient *models.Patient) error
	FindByID(id uint) (*models.Patient, error)
	Update(patient *models.Patient) error
	Delete(id uint) error
	FindAll() ([]models.Patient, error)
}
