package handlers

import (
	"tutorial/models"
	"tutorial/service"
)

type ChatSessionHandler struct {
	*stringCRUDHandler[models.ChatSession]
}

func NewChatSessionHandler(service service.ChatSessionServiceInterface) *ChatSessionHandler {
	return &ChatSessionHandler{stringCRUDHandler: newStringCRUDHandler[models.ChatSession](service)}
}
