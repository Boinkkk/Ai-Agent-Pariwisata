package handlers

import (
	"tutorial/models"
	"tutorial/service"
)

type PaymentHandler struct {
	*stringCRUDHandler[models.Payment]
}

func NewPaymentHandler(service service.PaymentServiceInterface) *PaymentHandler {
	return &PaymentHandler{stringCRUDHandler: newStringCRUDHandler[models.Payment](service)}
}
