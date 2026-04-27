package handlers

import (
	"tutorial/models"
	"tutorial/service"
)

type OrderStatusHistoryHandler struct {
	*stringCRUDHandler[models.OrderStatusHistory]
}

func NewOrderStatusHistoryHandler(service service.OrderStatusHistoryServiceInterface) *OrderStatusHistoryHandler {
	return &OrderStatusHistoryHandler{stringCRUDHandler: newStringCRUDHandler[models.OrderStatusHistory](service)}
}
