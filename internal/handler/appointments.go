package handler

import (
	"clinic/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handlers) createAppointment(c *gin.Context) {
	var input model.Appointment

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Appointment.Add(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
