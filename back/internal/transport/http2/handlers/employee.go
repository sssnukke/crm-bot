// handlers/employee.go
package handlers

import (
	"back/internal/dto"
	"back/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
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

	employeeData := &service.CreateEmployeeDto{
		Name:      req.Employee.Name,
		LastName:  req.Employee.LastName,
		Surname:   req.Employee.Surname,
		Phone:     req.Employee.Phone,
		Photo:     req.Employee.Photo,
		BirthDate: req.Employee.BirthDate,
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

func (h *EmployeeHandler) GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	employee, err := h.service.GetEmployeeById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
}

func (h *EmployeeHandler) DeleteEmployeeById(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	employee, err := h.service.DeleteById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
}

func (h *EmployeeHandler) UpdateEmployeePartial(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	employee, err := h.service.UpdateEmployeePartial(int64(id), updates)
	if err != nil {
		if err.Error() == "employee not found" {
			http.Error(w, "Employee not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
}
