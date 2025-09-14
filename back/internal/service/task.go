package service

import (
	"back/internal/models"
	"back/internal/repository"
	"errors"
	"time"
)

type TaskService struct {
	taskRepo     *repository.TaskRepository
	employeeRepo *repository.EmployeeRepository
}

func NewTaskService(taskRepo *repository.TaskRepository, employeeRepo *repository.EmployeeRepository) *TaskService {
	return &TaskService{
		taskRepo:     taskRepo,
		employeeRepo: employeeRepo,
	}
}

type TaskCreateDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	Status      string `json:"status"`
}

func (s *TaskService) CreateTask(employeeID int64, taskData *TaskCreateDTO) (*models.Task, error) {
	employee, err := s.employeeRepo.GetById(employeeID)
	if err != nil {
		return nil, err
	}
	if employee == nil {
		return nil, errors.New("employee not found")
	}

	deadline, err := time.Parse("2006-01-02", taskData.Deadline)
	if err != nil {
		return nil, errors.New("invalid deadline format, use YYYY-MM-DD")
	}

	task := &models.Task{
		Name:        taskData.Name,
		Description: taskData.Description,
		Deadline:    deadline,
		Status:      taskData.Status,
		EmployeeID:  employee.ID,
	}

	err = s.taskRepo.Create(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) GetTaskByID(taskID int64) (*models.Task, error) {
	return s.taskRepo.GetByID(taskID)
}

func (s *TaskService) GetTasks(employeeID int64) ([]*models.Task, error) {
	return s.taskRepo.GetAllByEmployeeID(employeeID)
}

func (s *TaskService) DeleteTaskByID(taskID int64) (*models.Task, error) {
	return s.taskRepo.DeleteById(taskID)
}
