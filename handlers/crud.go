package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type stringCRUDService[T any] interface {
	Create(ctx context.Context, item *T) error
	GetByID(ctx context.Context, id string) (*T, error)
	GetAll(ctx context.Context) ([]T, error)
	Update(ctx context.Context, id string, item *T) error
	Delete(ctx context.Context, id string) error
}

type intCRUDService[T any] interface {
	Create(ctx context.Context, item *T) error
	GetByID(ctx context.Context, id int) (*T, error)
	GetAll(ctx context.Context) ([]T, error)
	Update(ctx context.Context, id int, item *T) error
	Delete(ctx context.Context, id int) error
}

type stringCRUDHandler[T any] struct {
	service stringCRUDService[T]
}

type intCRUDHandler[T any] struct {
	service intCRUDService[T]
}

func newStringCRUDHandler[T any](service stringCRUDService[T]) *stringCRUDHandler[T] {
	return &stringCRUDHandler[T]{service: service}
}

func newIntCRUDHandler[T any](service intCRUDService[T]) *intCRUDHandler[T] {
	return &intCRUDHandler[T]{service: service}
}

func (h *stringCRUDHandler[T]) Create(c *gin.Context) {
	var item T
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.service.Create(ctx, &item); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (h *stringCRUDHandler[T]) FindByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	item, err := h.service.GetByID(ctx, c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *stringCRUDHandler[T]) FindAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	items, err := h.service.GetAll(ctx)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *stringCRUDHandler[T]) Update(c *gin.Context) {
	var item T
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.service.Update(ctx, c.Param("id"), &item); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "resource updated"})
}

func (h *stringCRUDHandler[T]) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.service.Delete(ctx, c.Param("id")); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "resource deleted"})
}

func (h *intCRUDHandler[T]) Create(c *gin.Context) {
	var item T
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.service.Create(ctx, &item); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (h *intCRUDHandler[T]) FindByID(c *gin.Context) {
	id, ok := parseIntParam(c, "id")
	if !ok {
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	item, err := h.service.GetByID(ctx, id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *intCRUDHandler[T]) FindAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	items, err := h.service.GetAll(ctx)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *intCRUDHandler[T]) Update(c *gin.Context) {
	id, ok := parseIntParam(c, "id")
	if !ok {
		return
	}

	var item T
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.service.Update(ctx, id, &item); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "resource updated"})
}

func (h *intCRUDHandler[T]) Delete(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"message": "resource deleted"})
}
