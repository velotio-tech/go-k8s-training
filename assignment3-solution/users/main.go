package main

import (

	//c "usersBackend/controllers"

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

	router.Run("localhost:9090")
}
