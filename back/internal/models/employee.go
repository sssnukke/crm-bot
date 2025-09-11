package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name      string `json:"name"`
	LastName  string `json:"lastName"`
	Surname   string `json:"surName"`
	Phone     string `json:"phone"`
	PhotoURL  string `json:"photoUrl"`
	BirthDate string `json:"birthDate"`
	UserId    uint
	Task      []Task `json:"tasks"`
}
