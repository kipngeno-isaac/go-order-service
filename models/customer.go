package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Id       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
}
