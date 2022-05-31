package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handlers) createAppointment(c *gin.Context) {
	// get doctorId, clientId and slot from query params

	c.JSON(http.StatusOK, gin.H{})
}
