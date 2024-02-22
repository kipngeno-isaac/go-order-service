package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kipngeno-isaac/go-order-service/controllers"
)

func RouteOrder(router *gin.Engine) {
	router.GET("/orders", controllers.GetOrders)
	router.POST("/orders/:customerId", controllers.CreateOrder)
	router.GET("/orders/:customerId", controllers.GetCustomerOrders)

}
