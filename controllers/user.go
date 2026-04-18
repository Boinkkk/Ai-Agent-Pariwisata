package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"tutorial/models"
	"tutorial/repository"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandle(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := h.repo.CreateUser(ctx, &user); err != nil {
		fmt.Println("Error Insert: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal Insert User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user berhasil dibuat",
		"user":    user,
	})
}

func (h *UserHandler) GetUserByName(c *gin.Context) {
	var user models.User

}
