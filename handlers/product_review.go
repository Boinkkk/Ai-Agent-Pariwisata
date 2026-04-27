package handlers

import (
	"tutorial/models"
	"tutorial/service"
)

type ProductReviewHandler struct {
	*stringCRUDHandler[models.ProductReview]
}

func NewProductReviewHandler(service service.ProductReviewServiceInterface) *ProductReviewHandler {
	return &ProductReviewHandler{stringCRUDHandler: newStringCRUDHandler[models.ProductReview](service)}
}
