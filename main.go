package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kipngeno-isaac/go-order-service/initializers"
	"github.com/kipngeno-isaac/go-order-service/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	routes.RouteCustomer(r)
	routes.RouteOrder(r)
	r.Run() // listen and serve on 0.0.0.0:8080}
}
