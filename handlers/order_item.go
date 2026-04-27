package handlers

import (
	"tutorial/models"
	"tutorial/service"
)

type OrderItemHandler struct {
	*stringCRUDHandler[models.OrderItem]
}

func NewOrderItemHandler(service service.OrderItemServiceInterface) *OrderItemHandler {
	return &OrderItemHandler{stringCRUDHandler: newStringCRUDHandler[models.OrderItem](service)}
}
