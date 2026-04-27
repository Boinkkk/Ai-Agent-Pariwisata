package handlers

import (
	"tutorial/models"
	"tutorial/service"
)

type CourierHandler struct {
	*intCRUDHandler[models.Courier]
}

func NewCourierHandler(service service.CourierServiceInterface) *CourierHandler {
	return &CourierHandler{intCRUDHandler: newIntCRUDHandler[models.Courier](service)}
}
