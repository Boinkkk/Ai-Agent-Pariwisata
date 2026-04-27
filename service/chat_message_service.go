package service

import (
	"context"
	"tutorial/models"
	"tutorial/repository"
)

type ChatMessageServiceInterface interface {
	Create(ctx context.Context, message *models.ChatMessage) error
	GetByID(ctx context.Context, id string) (*models.ChatMessage, error)
	GetAll(ctx context.Context) ([]models.ChatMessage, error)
	Update(ctx context.Context, id string, message *models.ChatMessage) error
	Delete(ctx context.Context, id string) error
}

type ChatMessageService struct {
	repo repository.ChatMessageRepositoryInterface
}

func NewChatMessageService(repo repository.ChatMessageRepositoryInterface) *ChatMessageService {
	return &ChatMessageService{repo: repo}
}
func (s *ChatMessageService) Create(ctx context.Context, message *models.ChatMessage) error {
	return s.repo.Insert(ctx, message)
}
func (s *ChatMessageService) GetByID(ctx context.Context, id string) (*models.ChatMessage, error) {
	return s.repo.FindByID(ctx, id)
}
func (s *ChatMessageService) GetAll(ctx context.Context) ([]models.ChatMessage, error) {
	return s.repo.FindAll(ctx)
}
func (s *ChatMessageService) Update(ctx context.Context, id string, message *models.ChatMessage) error {
	return s.repo.Update(ctx, id, message)
}
func (s *ChatMessageService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
