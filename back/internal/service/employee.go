// service/employee.go
package service

import (
	"back/internal/models"
	"back/internal/repository"
	"back/internal/utils"
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
	Name      string `json:"name"`
	LastName  string `json:"lastName"`
	Surname   string `json:"surName"`
	Phone     string `json:"phone"`
	Photo     string `json:"photo"`
	BirthDate string `json:"birthDate"`
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
		Name:      employeeData.Name,
		LastName:  employeeData.LastName,
		Surname:   employeeData.Surname,
		Phone:     employeeData.Phone,
		PhotoURL:  photoURL,
		BirthDate: employeeData.BirthDate,
		UserId:    user.ID,
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

func (s *EmployeeService) UpdateEmployeePartial(employeeID int64, updates map[string]interface{}) (*models.Employee, error) {
	employee, err := s.employeeRepo.GetById(employeeID)
	if err != nil {
		return nil, err
	}
	if employee == nil {
		return nil, errors.New("employee not found")
	}

	if photo, ok := updates["photo"].(string); ok && photo != "" {
		filename, err := utils.SaveBase64Image(photo, s.uploadDir)
		if err != nil {
			return nil, fmt.Errorf("failed to save photo: %v", err)
		}
		updates["photo_url"] = "/uploads/" + filename

		if employee.PhotoURL != "" {
			oldFilename := filepath.Base(employee.PhotoURL)
			os.Remove(filepath.Join(s.uploadDir, oldFilename))
		}

		delete(updates, "photo")
	}

	err = s.employeeRepo.UpdatePartial(employeeID, updates)
	if err != nil {
		return nil, err
	}

	return s.employeeRepo.GetById(employeeID)
}
