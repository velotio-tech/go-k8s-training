package main

import (
	"crypto/rand"
	"fmt"
	"ordersBackend/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	p, _ := rand.Prime(rand.Reader, 8)
	fmt.Println(p)

	router := gin.Default()

	router.POST("/user/:username/orders", controllers.CreateOrders)
	router.GET("/user/:username/orders", controllers.GetAllOrders)
	router.GET("/user/:username/orders/:order_id", controllers.GetOrderByOrderId)
	router.PUT("/user/:username/orders/:order_id/", controllers.UpdateOrder)
	router.DELETE("/user/:username/orders", controllers.DeleteAllOrders)
	router.DELETE("/user/:username/orders/:order_id", controllers.DeleteOrder)

	router.Run("localhost:9091")
}
