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

type loginInput struct {
	Name string `json:"name" binding:"required"`
}

func (h *Handler) login(c *gin.Context) {
	var input loginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.services.GenerateToken(input.Name)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}

func (h *Handler) logout(c *gin.Context) {}
