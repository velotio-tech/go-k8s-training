package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HeathCheck ...
func (h *UserHandlerClient) HeathCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"serverAlive":   true,
		"databaseAlive": h.domain.GetHealth(),
	})
}
