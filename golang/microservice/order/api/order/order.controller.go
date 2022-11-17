package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	AddItemHandler(*gin.Context)
	GetAllOrderHandler(*gin.Context)
	GetOrderHandler(*gin.Context)
	DeleteAllOrderHandler(*gin.Context)
	DeleteOrderHandler(*gin.Context)
}

type orderController struct {
	orderService OrderService
}

func NewOrderController(orderService OrderService) OrderController {
	return &orderController{
		orderService: orderService,
	}
}

func (oc orderController) AddItemHandler(ctx *gin.Context) {
	var addItemDTO AddItemDTO
	if err := ctx.Bind(&addItemDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := oc.orderService.AddItem(ctx.Request.Context(), addItemDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}

func (oc orderController) GetAllOrderHandler(ctx *gin.Context) {
	var orderAllDTO OrderAllDTO
	if err := ctx.Bind(&orderAllDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := oc.orderService.GetAllOrder(ctx.Request.Context(), orderAllDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}

func (oc orderController) GetOrderHandler(ctx *gin.Context) {
	var orderDTO OrderDTO
	if err := ctx.Bind(&orderDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := oc.orderService.GetOrder(ctx.Request.Context(), orderDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}

func (oc orderController) DeleteAllOrderHandler(ctx *gin.Context) {
	var orderAllDTO OrderAllDTO
	if err := ctx.Bind(&orderAllDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := oc.orderService.DeleteAllOrder(ctx.Request.Context(), orderAllDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}

func (oc orderController) DeleteOrderHandler(ctx *gin.Context) {
	var orderDTO OrderDTO
	if err := ctx.Bind(&orderDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := oc.orderService.DeleteOrder(ctx.Request.Context(), orderDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}
