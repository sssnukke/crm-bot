// service/employee.go
package service

import (
	"back/internal/models"
	"back/internal/repository"
	"back/internal/utils"
	"errors"
	"fmt"
)

type EmployeeService struct {
	employeeRepo *repository.EmployeeRepository
	userRepo     *repository.UserRepository
	uploadDir    string
}

func NewEmployeeService(employeeRepo *repository.EmployeeRepository, userRepo *repository.UserRepository, uploadDir string) *EmployeeService {
	return &EmployeeService{
		employeeRepo: employeeRepo,
		userRepo:     userRepo,
		uploadDir:    uploadDir,
	}
}

type CreateEmployeeDto struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Surname  string `json:"surName"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	Photo    string `json:"photo"`
}

func (s *EmployeeService) CreateEmployee(userId int64, employeeData *CreateEmployeeDto) (*models.Employee, error) {
	user, err := s.userRepo.GetByTgId(userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	var photoURL string
	if employeeData.Photo != "" {
		filename, err := utils.SaveBase64Image(employeeData.Photo, s.uploadDir)
		if err != nil {
			return nil, fmt.Errorf("failed to save photo: %v", err)
		}
		photoURL = "/uploads/" + filename
	}
	employee := &models.Employee{
		Name:     employeeData.Name,
		LastName: employeeData.LastName,
		Surname:  employeeData.Surname,
		Age:      employeeData.Age,
		Phone:    employeeData.Phone,
		PhotoURL: photoURL,
		UserId:   user.ID,
	}

	err = s.employeeRepo.Create(employee)
	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (s *EmployeeService) GetEmployeeById(id int64) (*models.Employee, error) {
	return s.employeeRepo.GetById(id)
}

func (s *EmployeeService) DeleteById(id int64) (*models.Employee, error) {
	return s.employeeRepo.DeleteById(id)
}
