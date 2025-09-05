package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Surname  string `json:"surName"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	UserId   uint
}
