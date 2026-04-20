package handlers

import (
	"context"
	"net/http"
	"time"
	"tutorial/repository"

	"github.com/gin-gonic/gin"
)

type CategoriesHandler struct {
	repo *repository.CategoriesRepository
}

func NewCategoriesHandler(repo *repository.CategoriesRepository) *CategoriesHandler {
	return &CategoriesHandler{repo: repo}
}

func (ctg *CategoriesHandler) GetCategories(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)

	defer cancel()

	categories, err := ctg.repo.GetCategories(ctx)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
	}

	c.JSON(http.StatusOK, categories)
}
