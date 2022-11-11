package main

import (
	"usersBackend/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/user", controllers.CreateUser)
	router.GET("/user/:username", controllers.GetUser)
	router.GET("/user", controllers.GetAllUsers)
	router.DELETE("/user/:username", controllers.DeleteUser)
	router.HEAD("/user/:username", controllers.GetUserMeta)

	router.POST("/user/:username/orders", controllers.CreateOrders)
	router.GET("/user/:username/orders", controllers.GetAllOrders)
	router.GET("/user/:username/orders/:order_id", controllers.GetOrderByOrderId)
	router.PUT("/user/:username/orders/:order_id", controllers.UpdateOrder)
	router.DELETE("/user/:username/orders", controllers.DeleteAllOrders)
	router.DELETE("/user/:username/orders/:order_id", controllers.DeleteOrder)

	router.Run("localhost:9090")
}
