package handlers

import (
	"context"
	"net/http"
	"time"
	"tutorial/models"
	"tutorial/service"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	service service.RoleServiceInterface
}

func NewRoleHandler(service service.RoleServiceInterface) *RoleHandler {
	return &RoleHandler{service: service}
}

func (h *RoleHandler) GetRole(c *gin.Context) {
	h.FindAll(c)
}

func (h *RoleHandler) Create(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.service.Create(ctx, &role); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, role)
}

func (h *RoleHandler) FindByID(c *gin.Context) {
	id, ok := parseIntParam(c, "id")
	if !ok {
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	role, err := h.service.GetByID(ctx, id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, role)
}

func (h *RoleHandler) FindAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	roles, err := h.service.GetAll(ctx)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, roles)
}

func (h *RoleHandler) Update(c *gin.Context) {
	id, ok := parseIntParam(c, "id")
	if !ok {
		return
	}

	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.service.Update(ctx, id, &role); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "role updated"})
}

func (h *RoleHandler) Delete(c *gin.Context) {
	id, ok := parseIntParam(c, "id")
	if !ok {
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.service.Delete(ctx, id); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "role deleted"})
}
