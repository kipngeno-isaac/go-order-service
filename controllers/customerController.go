package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kipngeno-isaac/go-order-service/initializers"
	"github.com/kipngeno-isaac/go-order-service/models"
)

func GetCustomers(c *gin.Context) {
	customers := []models.Customer{}

	initializers.DB.Find(&customers)
	c.JSON(200, &customers)
}

func CreateCustomer(c *gin.Context) {
	// get data from request body
	var CustomerData struct {
		Username string
		Phone    string
	}
	c.Bind(&CustomerData)
	// save data to db
	customer := models.Customer{
		Username: CustomerData.Username,
		Phone:    CustomerData.Phone,
	}
	result := initializers.DB.Create(&customer)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{"customer": customer})
}

func GetCustomer(c *gin.Context) {
	// get id fro url
	id := c.Param("id")

	// query customer with id
	var customer models.Customer
	initializers.DB.First(&customer, id)

	c.JSON(200, &customer)
}
