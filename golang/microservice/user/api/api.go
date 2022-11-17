package api

import (
	"github.com/gin-gonic/gin"
	"github.com/velotio-ajaykumbhar/microservice/user/api/auth"
	"github.com/velotio-ajaykumbhar/microservice/user/api/order"
	"github.com/velotio-ajaykumbhar/microservice/user/database"
)

func Setup() *gin.Engine {

	ds, err := database.Init()
	if err != nil {
		panic(err)
	}

	return InitRouter(ds)
}

func initAuthModule(ds *database.DataSource) auth.AuthController {
	service := auth.NewAuthService(ds)
	return auth.NewAuthController(service)
}

func InitRouter(ds *database.DataSource) *gin.Engine {
	router := gin.Default()

	authController := initAuthModule(ds)
	orderMicroService := order.NewOrderMicroService()

	router.POST("/register", authController.RegisterHandler)
	router.POST("/login", authController.LoginHandler)

	router.POST("/user/:userId/order/:orderId/item/:itemId", orderMicroService.AddItemHandler)

	router.GET("/user/:userId/order/:orderId", orderMicroService.GetOrderHandler)
	router.GET("/user/:userId/order", orderMicroService.GetAllOrderHandler)

	router.DELETE("/user/:userId/order/:orderId", orderMicroService.DeleteOrderHandler)
	router.DELETE("/user/:userId/order", orderMicroService.DeleteAllOrderHandler)

	router.GET("/user", authController.ReadAllHandler)

	return router
}
