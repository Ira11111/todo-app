package handlers

import (
	"github.com/Ira11111/todo-app/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) register(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Register(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//c.JSON(http.StatusOK, gin.H{"id": id})
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) login(c *gin.Context) {}

func (h *Handler) logout(c *gin.Context) {}
