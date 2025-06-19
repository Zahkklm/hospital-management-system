package repository

import (
	"database/sql"
	"hospital-management-system/internal/domain/models"
	"hospital-management-system/internal/domain/repository"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(user *models.User) error {
	query := `INSERT INTO users (username, password, role, created_at, updated_at) 
              VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id`

	err := r.db.QueryRow(query, user.Username, user.Password, user.Role).Scan(&user.ID)
	return err
}

func (r *UserRepositoryImpl) FindByID(id int) (*models.User, error) {
	query := `SELECT id, username, password, role, created_at, updated_at FROM users WHERE id = $1`

	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindByUsername(username string) (*models.User, error) {
	query := `SELECT id, username, password, role, created_at, updated_at FROM users WHERE username = $1`

	user := &models.User{}
	err := r.db.QueryRow(query, username).Scan(
		&user.ID, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) Update(user *models.User) error {
	query := `UPDATE users SET username = $1, password = $2, role = $3, updated_at = NOW() WHERE id = $4`

	_, err := r.db.Exec(query, user.Username, user.Password, user.Role, user.ID)
	return err
}

func (r *UserRepositoryImpl) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UserRepositoryImpl) FindAll() ([]models.User, error) {
	query := `SELECT id, username, password, role, created_at, updated_at FROM users ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
