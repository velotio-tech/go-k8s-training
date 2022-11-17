package api

import (
	"github.com/gin-gonic/gin"
	"github.com/velotio-ajaykumbhar/microservice/order/api/order"
	"github.com/velotio-ajaykumbhar/microservice/order/database"
)

func Setup() *gin.Engine {

	ds, err := database.Init()
	if err != nil {
		panic(err)
	}

	return InitRouter(ds)
}

func initOrderModule(ds *database.DataSource) order.OrderController {
	service := order.NewOrderService(ds)
	return order.NewOrderController(service)

}

func InitRouter(ds *database.DataSource) *gin.Engine {
	router := gin.Default()

	orderCotroller := initOrderModule(ds)

	router.POST("/addOrderItem", orderCotroller.AddItemHandler)

	router.POST("/getAllOrder", orderCotroller.GetAllOrderHandler)
	router.POST("/getOrder", orderCotroller.GetOrderHandler)

	router.POST("/deleteAllOrder", orderCotroller.DeleteAllOrderHandler)
	router.POST("/deleteOrder", orderCotroller.DeleteOrderHandler)

	return router
}
