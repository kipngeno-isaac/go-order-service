package models

type Customer struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
}
