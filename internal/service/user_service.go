package service

import (
	"user-management/internal/models"
	"user-management/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) RegisterUser(user *models.User) error {
	return s.repo.CreateUser(user)
}
