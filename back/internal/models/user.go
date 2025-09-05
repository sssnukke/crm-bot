package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	TgId     int64      `json:"tgId" gorm:"unique:not null"`
	Employee []Employee `json:"employees"`
}
