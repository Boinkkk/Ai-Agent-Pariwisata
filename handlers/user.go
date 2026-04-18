package handlers

import (
	"context"
	"net/http"
	"time"
	"tutorial/models"
	"tutorial/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	if err := h.service.CreateUser(ctx, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal Insert User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user berhasil dibuat",
		"user":    user,
	})
}
