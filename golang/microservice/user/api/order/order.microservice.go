package order

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderMicroService interface {
	AddItemHandler(*gin.Context)
	GetAllOrderHandler(*gin.Context)
	GetOrderHandler(*gin.Context)
	DeleteAllOrderHandler(*gin.Context)
	DeleteOrderHandler(*gin.Context)
}
type orderMicroService struct{}

func NewOrderMicroService() OrderMicroService {
	return &orderMicroService{}
}

func (oms orderMicroService) AddItemHandler(ctx *gin.Context) {
	var addItemDTO AddItemDTO
	if err := ctx.BindUri(&addItemDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var body bytes.Buffer
	orderRequest := OrderRequest{
		OrderId: addItemDTO.OrderId,
		UserId:  addItemDTO.UserId,
		Item:    addItemDTO.Item,
	}

	err := json.NewEncoder(&body).Encode(orderRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := http.Post("http://order-app:9090/addOrderItem", "application/json", &body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer response.Body.Close()

	var responseBody map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responseBody)

}

func (oms orderMicroService) GetAllOrderHandler(ctx *gin.Context) {
	log.Println("[order.miscroservice] getAllOrderHandler")
	var orderAllDTO OrderAllDTO
	if err := ctx.BindUri(&orderAllDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var body bytes.Buffer
	orderRequest := OrderRequest{
		UserId: orderAllDTO.UserId,
	}

	err := json.NewEncoder(&body).Encode(orderRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := http.Post("http://order-app:9090/getAllOrder", "application/json", &body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer response.Body.Close()

	var responseBody map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responseBody)
}

func (oms orderMicroService) GetOrderHandler(ctx *gin.Context) {
	log.Println("[order.miscroservice] getOrderHandler")
	var orderDTO OrderDTO
	if err := ctx.BindUri(&orderDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var body bytes.Buffer
	orderRequest := OrderRequest{
		UserId:  orderDTO.UserId,
		OrderId: orderDTO.OrderId,
	}

	err := json.NewEncoder(&body).Encode(orderRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := http.Post("http://order-app:9090/getOrder", "application/json", &body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer response.Body.Close()

	var responseBody map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responseBody)
}

func (oms orderMicroService) DeleteAllOrderHandler(ctx *gin.Context) {
	log.Println("[order.miscroservice] getAllOrderHandler")
	var orderAllDTO OrderAllDTO
	if err := ctx.BindUri(&orderAllDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var body bytes.Buffer
	orderRequest := OrderRequest{
		UserId: orderAllDTO.UserId,
	}

	err := json.NewEncoder(&body).Encode(orderRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := http.Post("http://order-app:9090/deleteAllOrder", "application/json", &body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer response.Body.Close()

	var responseBody map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responseBody)
}

func (oms orderMicroService) DeleteOrderHandler(ctx *gin.Context) {
	log.Println("[order.miscroservice] getOrderHandler")
	var orderDTO OrderDTO
	if err := ctx.BindUri(&orderDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var body bytes.Buffer
	orderRequest := OrderRequest{
		UserId:  orderDTO.UserId,
		OrderId: orderDTO.OrderId,
	}

	err := json.NewEncoder(&body).Encode(orderRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := http.Post("http://order-app:9090/deleteOrder", "application/json", &body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer response.Body.Close()

	var responseBody map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responseBody)
}
