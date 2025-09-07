package repository

import (
	"back/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByTgId(tgId int64) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Employee").Where("tg_id = ?", tgId).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
