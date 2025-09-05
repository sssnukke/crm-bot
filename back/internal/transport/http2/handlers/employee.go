// handlers/employee.go
package handlers

import (
	"back/internal/dto"
	"back/internal/service"
	"encoding/json"
	"net/http"
)

type EmployeeHandler struct {
	service *service.EmployeeService
}

func NewEmployeeHandler(service *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateEmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	println("Received request - UserId:", req.UserId)
	println("Employee data - Name:", req.Employee.Name)
	println("Employee data - LastName:", req.Employee.LastName)

	employeeData := &service.CreateEmployeeDto{
		Name:     req.Employee.Name,
		LastName: req.Employee.LastName,
		Surname:  req.Employee.Surname,
		Age:      req.Employee.Age,
		Phone:    req.Employee.Phone,
	}

	employee, err := h.service.CreateEmployee(req.UserId, employeeData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employee)
}
