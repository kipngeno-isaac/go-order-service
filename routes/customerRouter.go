package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kipngeno-isaac/go-order-service/controllers"
)

func RouteCustomer(router *gin.Engine) {
	router.GET("/customers", controllers.GetCustomers)
	router.POST("/customers", controllers.CreateCustomer)
	router.GET("/customers/:id", controllers.GetCustomer)

}
