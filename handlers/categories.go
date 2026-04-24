package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"tutorial/models"
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

func (ctg *CategoriesHandler) AddCategories(c *gin.Context) {

	var request *models.Categories

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request",
		})

		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	err = ctg.repo.AddCategorie(ctx, request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Request Invalid",
			"Message": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succes",
		"add": gin.H{
			"name":        request.Name,
			"slug":        request.Slug,
			"description": request.Description,
		},
	})

}

func (ctg *CategoriesHandler) DeleteCategoriesByID(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	if err := ctg.repo.DeleteCategoriesByID(ctx, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "ID tidak tersedia",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Kategori berhasil dihapus",
	})
}
