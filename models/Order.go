package models

type Order struct {
	Id         int    `json:"id" gorm:"primary_key"`
	Item       string `json:"item"`
	Amount     int    `json:"amount"`
	CustomerId int    `json:"customer_id"`
}
