package handlers

import (
	"tutorial/models"
	"tutorial/service"
)

type ChatMessageHandler struct {
	*stringCRUDHandler[models.ChatMessage]
}

func NewChatMessageHandler(service service.ChatMessageServiceInterface) *ChatMessageHandler {
	return &ChatMessageHandler{stringCRUDHandler: newStringCRUDHandler[models.ChatMessage](service)}
}
