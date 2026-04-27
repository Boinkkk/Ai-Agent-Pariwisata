package handlers

import (
	"tutorial/models"
	"tutorial/service"
)

type CartHandler struct {
	*stringCRUDHandler[models.Cart]
}

func NewCartHandler(service service.CartServiceInterface) *CartHandler {
	return &CartHandler{stringCRUDHandler: newStringCRUDHandler[models.Cart](service)}
}
