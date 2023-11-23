package handlers

import (
	"net/http"

	"ums/pkg/exception"
	"ums/pkg/models"

	"github.com/gin-gonic/gin"
)

// UserOrderCreateRequest ...
type UserOrderCreateRequest struct {
	BuyerUserID string          `json:"buyer_user_id" binding:"required"`
	Status      string          `json:"status" binding:"required"`
	TotalAmount uint64          `json:"total_amount" binding:"required"`
	Products    models.Products `json:"products" binding:"required"`
	PaymentID   string          `json:"payment_id" binding:"required"`
}

// CreateOrder ...
func (h *UserHandlerClient) CreateOrder(c *gin.Context) {
	rootTracer := h.traceEveryRequest(c)
	defer rootTracer.Finish()

	userID, exits := c.Params.Get("userID")
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id is required"})
		return
	}

	payload := &UserOrderCreateRequest{}
	exceptionObj := h.hlpr.DecodeJSONBody(c.Request, &payload)
	if exceptionObj != nil {
		h.serveRes(c, exceptionObj.StatusCode, map[string]interface{}{"error": exceptionObj}, &exception.Exception{
			Err:        exceptionObj.Err,
			Message:    exceptionObj.Message,
			StatusCode: exceptionObj.StatusCode,
			StatusText: exceptionObj.StatusText,
		}, rootTracer)
		return
	}

	_, errObj := h.domain.GetUser(userID, "")
	if errObj != nil {
		h.serveRes(c, errObj.StatusCode, map[string]interface{}{"error": errObj}, errObj, rootTracer)
		return
	}

	orderID, errObj := h.service.OrderService.CreateOrder(&models.Order{
		BuyerUserID: payload.BuyerUserID,
		Status:      payload.Status,
		PaymentID:   payload.PaymentID,
		Products:    payload.Products,
	})
	if errObj != nil {
		h.serveRes(c, errObj.StatusCode, map[string]interface{}{"error": errObj}, errObj, rootTracer)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"orderID": orderID,
	})
}
