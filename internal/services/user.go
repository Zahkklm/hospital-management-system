package services

import (
    "hospital-management-system/internal/domain/models"
    "hospital-management-system/internal/domain/repository"
)

type UserService struct {
    userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
    return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
    return s.userRepo.FindByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
    return s.userRepo.Update(user)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
    return s.userRepo.FindAll()
}