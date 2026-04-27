package service

import (
	"context"
	"tutorial/models"
	"tutorial/repository"
)

type ChatSessionServiceInterface interface {
	Create(ctx context.Context, session *models.ChatSession) error
	GetByID(ctx context.Context, id string) (*models.ChatSession, error)
	GetAll(ctx context.Context) ([]models.ChatSession, error)
	Update(ctx context.Context, id string, session *models.ChatSession) error
	Delete(ctx context.Context, id string) error
}

type ChatSessionService struct {
	repo repository.ChatSessionRepositoryInterface
}

func NewChatSessionService(repo repository.ChatSessionRepositoryInterface) *ChatSessionService {
	return &ChatSessionService{repo: repo}
}
func (s *ChatSessionService) Create(ctx context.Context, session *models.ChatSession) error {
	return s.repo.Insert(ctx, session)
}
func (s *ChatSessionService) GetByID(ctx context.Context, id string) (*models.ChatSession, error) {
	return s.repo.FindByID(ctx, id)
}
func (s *ChatSessionService) GetAll(ctx context.Context) ([]models.ChatSession, error) {
	return s.repo.FindAll(ctx)
}
func (s *ChatSessionService) Update(ctx context.Context, id string, session *models.ChatSession) error {
	return s.repo.Update(ctx, id, session)
}
func (s *ChatSessionService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
