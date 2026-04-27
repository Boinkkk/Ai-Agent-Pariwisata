package handlers

import (
	"tutorial/models"
	"tutorial/service"
)

type ArticleHandler struct {
	*stringCRUDHandler[models.Article]
}

func NewArticleHandler(service service.ArticleServiceInterface) *ArticleHandler {
	return &ArticleHandler{stringCRUDHandler: newStringCRUDHandler[models.Article](service)}
}
