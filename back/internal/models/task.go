package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Status      string    `json:"status"`
	EmployeeID  uint
}
