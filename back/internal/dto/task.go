package dto

type CreateTaskRequest struct {
	EmployeeID int64          `json:"employeeId"`
	Task       CreateTaskData `json:"task"`
}

type CreateTaskData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	Status      string `json:"status"`
}
