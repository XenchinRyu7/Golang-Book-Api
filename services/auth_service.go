package services

import (
	"errors"
	"github.com/google/uuid"
	"golang-book-api/models"
	"golang-book-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(username, password, email string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       uuid.New(),
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(username, password string) (*models.User, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}
