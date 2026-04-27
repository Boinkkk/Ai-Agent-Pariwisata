package handlers

import (
	"tutorial/models"
	"tutorial/service"
)

type KnowledgeBaseHandler struct {
	*stringCRUDHandler[models.KnowledgeBase]
}

func NewKnowledgeBaseHandler(service service.KnowledgeBaseServiceInterface) *KnowledgeBaseHandler {
	return &KnowledgeBaseHandler{stringCRUDHandler: newStringCRUDHandler[models.KnowledgeBase](service)}
}
