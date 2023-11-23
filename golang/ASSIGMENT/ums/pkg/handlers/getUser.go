package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUser ...
func (h *UserHandlerClient) GetUser(c *gin.Context) {
	rootTracer := h.traceEveryRequest(c)
	defer rootTracer.Finish()

	userID, exits := c.Params.Get("userID")
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id is required"})
		return
	}

	user, errObj := h.domain.GetUser(userID, "")
	if errObj != nil {
		h.serveRes(c, errObj.StatusCode, map[string]interface{}{"error": errObj}, errObj, rootTracer)
		return
	}

	h.serveRes(c, http.StatusOK, map[string]interface{}{
		"user": user,
	}, nil, rootTracer)
}

// GetUsers ...
func (h *UserHandlerClient) GetUsers(c *gin.Context) {
	rootTracer := h.traceEveryRequest(c)
	defer rootTracer.Finish()

	users, errObj := h.domain.GetUsers()
	if errObj != nil {
		h.serveRes(c, errObj.StatusCode, map[string]interface{}{"error": errObj}, errObj, rootTracer)
		return
	}

	h.serveRes(c, http.StatusOK, map[string]interface{}{
		"users": users,
	}, nil, rootTracer)
}
