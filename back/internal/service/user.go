package service

import (
	"back/internal/models"
	"back/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(tgId int64) (*models.User, error) {
	user := &models.User{TgId: tgId}
	err := s.repo.Create(user)
	return user, err
}
