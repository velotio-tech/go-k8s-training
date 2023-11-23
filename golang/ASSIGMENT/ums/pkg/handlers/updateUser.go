package handlers

import (
	"net/http"

	"ums/pkg/exception"

	"github.com/gin-gonic/gin"
)

type UserUpdateRequest struct {
	Name string `json:"name"`
}

// UpdateUser ...
func (h *UserHandlerClient) UpdateUser(c *gin.Context) {
	rootTracer := h.traceEveryRequest(c)
	defer rootTracer.Finish()

	userID, exits := c.Params.Get("userID")
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id is required"})
		return
	}

	payload := &UserUpdateRequest{}
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

	errObj := h.domain.UpdateUser(payload.Name, userID)
	if errObj != nil {
		h.serveRes(c, errObj.StatusCode, map[string]interface{}{"error": errObj}, errObj, rootTracer)
		return
	}

	h.serveRes(c, http.StatusOK, map[string]interface{}{
		"message": "user details updated successfully",
	}, nil, rootTracer)
}
