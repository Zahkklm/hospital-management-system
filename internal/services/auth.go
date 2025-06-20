package services

import (
	"errors"
	"hospital-management-system/internal/domain/models"
	"hospital-management-system/internal/domain/repository"
	"hospital-management-system/pkg/utils"
)

type AuthService struct {
	userRepo repository.UserRepository
	secret   string
}

func NewAuthService(userRepo repository.UserRepository, secret string) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		secret:   secret,
	}
}

func (s *AuthService) Register(user *models.User) error {
	// Check if username already exists
	existingUser, _ := s.userRepo.FindByUsername(user.Username)
	if existingUser != nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.userRepo.Create(user)
}

// Original Login method (for backward compatibility)
func (s *AuthService) Login(username, password string) (string, error) {
	_, token, err := s.LoginWithUser(username, password)
	if err != nil {
		return "", err
	}
	return token, nil
}

// New method that returns both user and token
func (s *AuthService) LoginWithUser(username, password string) (*models.User, string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, "", errors.New("invalid credentials")
	}

	// Generate token using the available function
	token, err := utils.GenerateToken(user.Username, user.Role)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *AuthService) ValidateToken(token string) (*models.User, error) {
	claims, err := utils.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	// Create user object with token data
	user := &models.User{
		Username: claims.Username,
		Role:     claims.Role,
	}

	return user, nil
}
