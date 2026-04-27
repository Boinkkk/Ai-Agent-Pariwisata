package handlers

import (
	"tutorial/models"
	"tutorial/service"
)

type AddressesHandler struct {
	*stringCRUDHandler[models.Addresses]
}

func NewAddressesHandler(service service.AddressesServiceInterface) *AddressesHandler {
	return &AddressesHandler{stringCRUDHandler: newStringCRUDHandler[models.Addresses](service)}
}
