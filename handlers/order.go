package handlers

import (
	"tutorial/models"
	"tutorial/service"
)

type OrderHandler struct {
	*stringCRUDHandler[models.Order]
}

func NewOrderHandler(service service.OrderServiceInterface) *OrderHandler {
	return &OrderHandler{stringCRUDHandler: newStringCRUDHandler[models.Order](service)}
}
