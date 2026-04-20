package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"
	dto "tutorial/DTO"
	"tutorial/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Login Error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "input Invalid",
		})

		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	user, err := h.service.UserLogin(ctx, req.Email, req.Password)

	if err != nil {
		fmt.Println("INvalid email assword", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Email Or Password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Sucess",
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
		},
	})
}
