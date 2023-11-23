package handlers

import (
	"net/http"

	"ums/pkg/exception"
	"ums/pkg/models"

	"github.com/gin-gonic/gin"
)

// UserCreateRequest ...
type UserCreateRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

// CreateUser ...
func (h *UserHandlerClient) CreateUser(c *gin.Context) {
	rootTracer := h.traceEveryRequest(c)
	defer rootTracer.Finish()

	payload := &UserCreateRequest{}
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

	createdID, errObj := h.domain.CreateUser(&models.User{
		Name:  payload.Name,
		Email: payload.Email,
	})
	if errObj != nil {
		h.serveRes(c, exceptionObj.StatusCode, map[string]interface{}{"error": exceptionObj}, &exception.Exception{
			Err:        errObj.Err,
			Message:    errObj.Message,
			StatusCode: errObj.StatusCode,
			StatusText: errObj.StatusText,
		}, rootTracer)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"userID": createdID,
	})
}
