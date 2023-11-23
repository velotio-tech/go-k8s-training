package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteOrder ...
func (h *UserHandlerClient) DeleteOrder(c *gin.Context) {
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

	errObj = h.service.OrderService.DeleteOrder(orderID)
	if errObj != nil {
		h.serveRes(c, errObj.StatusCode, map[string]interface{}{"error": errObj}, errObj, rootTracer)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order deleted successfully",
	})
}

// DeleteOrders ...
func (h *UserHandlerClient) DeleteOrders(c *gin.Context) {
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

	errObj = h.service.OrderService.DeleteOrders(userID)
	if errObj != nil {
		h.serveRes(c, errObj.StatusCode, map[string]interface{}{"error": errObj}, errObj, rootTracer)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Orders deleted successfully",
	})
}
