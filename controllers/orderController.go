package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kipngeno-isaac/go-order-service/initializers"
	"github.com/kipngeno-isaac/go-order-service/models"
)

func GetOrders(c *gin.Context) {
	orders := []models.Order{}

	initializers.DB.Find(&orders)
	c.JSON(200, &orders)
}

func CreateOrder(c *gin.Context) {
	// get customer id
	customerId, _ := strconv.Atoi(c.Param("customerId"))

	// get data from request body
	var orderData struct {
		Item   string
		Amount int
	}
	c.Bind(&orderData)
	// save data to db
	order := models.Order{
		Id:         0,
		Item:       orderData.Item,
		Amount:     orderData.Amount,
		CustomerId: customerId,
	}
	result := initializers.DB.Create(&order)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// TODO: integrate with africaistalking api

	c.JSON(200, gin.H{"Order": order})
}

func GetCustomerOrders(c *gin.Context) {
	// get id fro url
	customerId := c.Param("customerId")

	// query orders with customer id
	var orders []models.Order
	initializers.DB.Where("customer_id", customerId).Find(&orders)

	c.JSON(200, &orders)
}
