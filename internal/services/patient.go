package services

import (
    "errors"
    "hospital-management-system/internal/domain/models"
    "hospital-management-system/internal/domain/repository"
)

type PatientService struct {
    repo repository.PatientRepository
}

func NewPatientService(repo repository.PatientRepository) *PatientService {
    return &PatientService{repo: repo}
}

func (s *PatientService) CreatePatient(patient *models.Patient) error {
    if patient.FirstName == "" || patient.LastName == "" {
        return errors.New("first name and last name are required")
    }
    
    if patient.Email == "" {
        return errors.New("email is required")
    }
    
    return s.repo.Create(patient)
}

func (s *PatientService) GetPatientByID(id uint) (*models.Patient, error) {
    patient, err := s.repo.FindByID(id)
    if err != nil {
        return nil, err
    }
    return patient, nil
}

func (s *PatientService) UpdatePatient(patient *models.Patient) error {
    existingPatient, err := s.repo.FindByID(uint(patient.ID))
    if err != nil {
        return err
    }
    if existingPatient == nil {
        return errors.New("patient not found")
    }
    return s.repo.Update(patient)
}

func (s *PatientService) DeletePatient(id uint) error {
    existingPatient, err := s.repo.FindByID(id)
    if err != nil {
        return err
    }
    if existingPatient == nil {
        return errors.New("patient not found")
    }
    return s.repo.Delete(id)
}

func (s *PatientService) GetAllPatients() ([]models.Patient, error) {
    patients, err := s.repo.FindAll()
    if err != nil {
        return nil, err
    }
    return patients, nil
}