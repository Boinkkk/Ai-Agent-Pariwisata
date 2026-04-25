package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
	"tutorial/logger"
	"tutorial/models"
	"tutorial/repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type ProductHandler struct {
	repo *repository.ProductRepository
}

func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (h *ProductHandler) GetAllProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)

	defer cancel()

	products, err := h.repo.GetAllProduct(ctx)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "error",
			"Message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)

	defer cancel()

	id := c.Param("id")

	product, err := h.repo.GetProductByID(ctx, id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {

			logger.Log.Warn("Product not found", zap.String("id", id), zap.Error(err))
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "Product not found",
			})
			return
		}

		logger.Log.Error("Failed to Get Product", zap.String("id", id), zap.Error(err))

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) AddProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		logger.Log.Warn("Body Request tidak valid", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Body Request tidak valid",
		})

		return
	}

	if err := h.repo.AddProduct(ctx, &product); err != nil {
		logger.Log.Error("Error adding Product", zap.Any("body", product), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Error menambahkan produk",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"product": product,
	})

}

func (h *ProductHandler) UpdateProductByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	id := c.Param("id")

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		logger.Log.Error("Failed Bind Product", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Field Request tidak sesuai",
		})
	}

	fmt.Println(product)

	err := h.repo.UpdateProductByID(ctx, id, product)

	if err != nil {
		if err == pgx.ErrNoRows {
			logger.Log.Warn("Product Not Updated", zap.String("id", id), zap.Error(err))

			c.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "Produk gagal di update (produk tidak ditemukan)",
			})

			return
		}

		logger.Log.Error("Failed to update product", zap.String("id", id), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"Message": "Gagal Update Produk",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "sucess",
		"message": "Produk berhasil di Update",
	})
}

func (h *ProductHandler) DeleteProductByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Param("id")
	err := h.repo.DeleteProductByID(ctx, id)

	if err != nil {
		if err == pgx.ErrNoRows {
			logger.Log.Warn("Produk Tidak ditemukan", zap.String("id", id), zap.Error(err))
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "Produk Gagal di delete (Produk tidak ditemukan)",
			})
			return
		}

		logger.Log.Error("Gagal melakukan delete", zap.String("id", id), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Tidak dapat melakukan delete produk",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "succes",
		"Message": "Produk Berhasil di delete",
	})

}
