package repository

import (
	"back/internal/models"
	"errors"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Create(employee *models.Employee) error {
	return r.db.Create(employee).Error
}

func (r *EmployeeRepository) UpdatePartial(id int64, updates map[string]interface{}) error {
	result := r.db.Model(&models.Employee{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *EmployeeRepository) GetById(id int64) (*models.Employee, error) {
	var employee models.Employee
	err := r.db.Where("id = ?", id).First(&employee).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &employee, nil
}

func (r *EmployeeRepository) DeleteById(id int64) (*models.Employee, error) {
	var employee models.Employee

	if err := r.db.First(&employee, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Unscoped().Delete(&employee).Error; err != nil {
		return nil, err
	}

	return &employee, nil
}
