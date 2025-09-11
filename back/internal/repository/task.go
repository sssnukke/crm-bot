package repository

import (
	"back/internal/models"
	"errors"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) DeleteById(id int64) (*models.Task, error) {
	var task models.Task

	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Unscoped().Delete(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) GetByID(id int64) (*models.Task, error) {
	var task models.Task

	err := r.db.Where("id = ?", id).First(&task).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) GetAllByEmployeeID(employeeID int64) ([]*models.Task, error) {
	var tasks []*models.Task

	err := r.db.Where("employee_id = ?", employeeID).Find(&tasks).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return tasks, nil
}
