package repository

import (
	"database/sql"
	"hospital-management-system/internal/domain/models"
	"hospital-management-system/internal/domain/repository"
)

type PatientRepositoryImpl struct {
	db *sql.DB
}

func NewPatientRepository(db *sql.DB) repository.PatientRepository {
	return &PatientRepositoryImpl{db: db}
}

func (r *PatientRepositoryImpl) Create(patient *models.Patient) error {
	query := `INSERT INTO patients (first_name, last_name, date_of_birth, gender, phone_number, email, address, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW()) RETURNING id`

	err := r.db.QueryRow(query, patient.FirstName, patient.LastName, patient.DOB,
		patient.Gender, patient.Phone, patient.Email, patient.Address).Scan(&patient.ID)

	return err
}

func (r *PatientRepositoryImpl) FindByID(id uint) (*models.Patient, error) {
	query := `SELECT id, first_name, last_name, date_of_birth, gender, phone_number, email, address, created_at, updated_at 
              FROM patients WHERE id = $1`

	patient := &models.Patient{}
	err := r.db.QueryRow(query, id).Scan(
		&patient.ID, &patient.FirstName, &patient.LastName, &patient.DOB,
		&patient.Gender, &patient.Phone, &patient.Email, &patient.Address,
		&patient.CreatedAt, &patient.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return patient, nil
}

func (r *PatientRepositoryImpl) Update(patient *models.Patient) error {
	query := `UPDATE patients SET first_name = $1, last_name = $2, date_of_birth = $3, 
              gender = $4, phone_number = $5, email = $6, address = $7, updated_at = NOW() 
              WHERE id = $8`

	_, err := r.db.Exec(query, patient.FirstName, patient.LastName, patient.DOB,
		patient.Gender, patient.Phone, patient.Email, patient.Address, patient.ID)

	return err
}

func (r *PatientRepositoryImpl) Delete(id uint) error {
	query := `DELETE FROM patients WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PatientRepositoryImpl) FindAll() ([]models.Patient, error) {
	query := `SELECT id, first_name, last_name, date_of_birth, gender, phone_number, email, address, created_at, updated_at 
              FROM patients ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []models.Patient
	for rows.Next() {
		var patient models.Patient
		err := rows.Scan(
			&patient.ID, &patient.FirstName, &patient.LastName, &patient.DOB,
			&patient.Gender, &patient.Phone, &patient.Email, &patient.Address,
			&patient.CreatedAt, &patient.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}

	return patients, nil
}
