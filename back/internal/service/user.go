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

func (s *UserService) CheckUserExists(tgId int64) (bool, error) {
	user, err := s.repo.GetByTgId(tgId)
	if err != nil {
		return false, err
	}

	return user != nil, nil
}

func (s *UserService) GetUserByTgId(tgId int64) (*models.User, error) {
	return s.repo.GetByTgId(tgId)
}
