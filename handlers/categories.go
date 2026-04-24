package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
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

func (ctg *CategoriesHandler) DeleteCategoriesByID(c *gin.Context) error {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Errorf("Invalid Input")
		return err
	}

	if err := ctg.repo.DeleteCategoriesByID(ctx, id); err != nil {
		fmt.Errorf("Id Tidak Tersedia")
		return err
	}

	return nil
}
