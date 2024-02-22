package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kipngeno-isaac/go-order-service/controllers"
)

func RouteCustomer(router *gin.Engine) {
	router.GET("/", controllers.GetCustomers)
	router.POST("/", controllers.CreateCustomer)
	router.GET("/:id", controllers.GetCustomer)

}
