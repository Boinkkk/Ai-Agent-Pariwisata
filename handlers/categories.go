package handlers

import (
	"context"
	"net/http"
	"time"
	"tutorial/models"
	"tutorial/service"

	"github.com/gin-gonic/gin"
)

type CategoriesHandler struct {
	service service.CategoriesServiceInterface
}

func NewCategoriesHandler(service service.CategoriesServiceInterface) *CategoriesHandler {
	return &CategoriesHandler{service: service}
}

func (h *CategoriesHandler) GetCategories(c *gin.Context) {
	h.FindAll(c)
}

func (h *CategoriesHandler) AddCategories(c *gin.Context) {
	h.Create(c)
}

func (h *CategoriesHandler) DeleteCategoriesByID(c *gin.Context) {
	h.Delete(c)
}

func (h *CategoriesHandler) Create(c *gin.Context) {
	var category models.Categories
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.service.Create(ctx, &category); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, category)
}

func (h *CategoriesHandler) FindByID(c *gin.Context) {
	id, ok := parseIntParam(c, "id")
	if !ok {
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	category, err := h.service.GetByID(ctx, id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *CategoriesHandler) FindAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	categories, err := h.service.GetAll(ctx)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (h *CategoriesHandler) Update(c *gin.Context) {
	id, ok := parseIntParam(c, "id")
	if !ok {
		return
	}

	var category models.Categories
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.service.Update(ctx, id, &category); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "category updated"})
}

func (h *CategoriesHandler) Delete(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"message": "category deleted"})
}
