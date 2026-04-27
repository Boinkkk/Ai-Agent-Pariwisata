package handlers

import (
	"tutorial/models"
	"tutorial/service"
)

type CartItemHandler struct {
	*stringCRUDHandler[models.CartItem]
}

func NewCartItemHandler(service service.CartItemServiceInterface) *CartItemHandler {
	return &CartItemHandler{stringCRUDHandler: newStringCRUDHandler[models.CartItem](service)}
}
