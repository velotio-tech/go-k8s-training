package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOrder ...
func (h *UserHandlerClient) GetOrder(c *gin.Context) {
	rootTracer := h.traceEveryRequest(c)
	defer rootTracer.Finish()

	userID, exits := c.Params.Get("userID")
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id is required"})
		return
	}

	orderID, exits := c.Params.Get("orderID")
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order id is required"})
		return
	}

	_, errObj := h.domain.GetUser(userID, "")
	if errObj != nil {
		h.serveRes(c, errObj.StatusCode, map[string]interface{}{"error": errObj}, errObj, rootTracer)
		return
	}

	order, errObj := h.service.OrderService.GetOrderByID(orderID)
	if errObj != nil {
		h.serveRes(c, errObj.StatusCode, map[string]interface{}{"error": errObj}, errObj, rootTracer)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order": order,
	})
}

// GetOrders ...
func (h *UserHandlerClient) GetOrders(c *gin.Context) {
	rootTracer := h.traceEveryRequest(c)
	defer rootTracer.Finish()

	userID, exits := c.Params.Get("userID")
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id is required"})
		return
	}

	_, errObj := h.domain.GetUser(userID, "")
	if errObj != nil {
		h.serveRes(c, errObj.StatusCode, map[string]interface{}{"error": errObj}, errObj, rootTracer)
		return
	}

	orders, errObj := h.service.OrderService.GetOrders(userID)
	if errObj != nil {
		h.serveRes(c, errObj.StatusCode, map[string]interface{}{"error": errObj}, errObj, rootTracer)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orders	": orders,
	})
}
