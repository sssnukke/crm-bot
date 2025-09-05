// service/employee.go
package service

import (
	"back/internal/models"
	"back/internal/repository"
	"errors"
)

type EmployeeService struct {
	employeeRepo *repository.EmployeeRepository
	userRepo     *repository.UserRepository
}

func NewEmployeeService(employeeRepo *repository.EmployeeRepository, userRepo *repository.UserRepository) *EmployeeService {
	return &EmployeeService{
		employeeRepo: employeeRepo,
		userRepo:     userRepo,
	}
}

type CreateEmployeeDto struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Surname  string `json:"surName"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
}

func (s *EmployeeService) CreateEmployee(userId int64, employeeData *CreateEmployeeDto) (*models.Employee, error) {
	// 1. Находим пользователя по ID
	user, err := s.userRepo.GetByTgId(userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// 2. Создаем сотрудника и привязываем к пользователю
	employee := &models.Employee{
		Name:     employeeData.Name,
		LastName: employeeData.LastName,
		Surname:  employeeData.Surname,
		Age:      employeeData.Age,
		Phone:    employeeData.Phone,
		UserId:   user.ID, // ← Привязываем к пользователю
	}

	// 3. Сохраняем сотрудника
	err = s.employeeRepo.Create(employee)
	if err != nil {
		return nil, err
	}

	return employee, nil
}
