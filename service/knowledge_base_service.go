package service

import (
	"context"
	"tutorial/models"
	"tutorial/repository"
)

type KnowledgeBaseServiceInterface interface {
	Create(ctx context.Context, knowledge *models.KnowledgeBase) error
	GetByID(ctx context.Context, id string) (*models.KnowledgeBase, error)
	GetAll(ctx context.Context) ([]models.KnowledgeBase, error)
	Update(ctx context.Context, id string, knowledge *models.KnowledgeBase) error
	Delete(ctx context.Context, id string) error
}

type KnowledgeBaseService struct {
	repo repository.KnowledgeBaseRepositoryInterface
}

func NewKnowledgeBaseService(repo repository.KnowledgeBaseRepositoryInterface) *KnowledgeBaseService {
	return &KnowledgeBaseService{repo: repo}
}
func (s *KnowledgeBaseService) Create(ctx context.Context, knowledge *models.KnowledgeBase) error {
	return s.repo.Insert(ctx, knowledge)
}
func (s *KnowledgeBaseService) GetByID(ctx context.Context, id string) (*models.KnowledgeBase, error) {
	return s.repo.FindByID(ctx, id)
}
func (s *KnowledgeBaseService) GetAll(ctx context.Context) ([]models.KnowledgeBase, error) {
	return s.repo.FindAll(ctx)
}
func (s *KnowledgeBaseService) Update(ctx context.Context, id string, knowledge *models.KnowledgeBase) error {
	return s.repo.Update(ctx, id, knowledge)
}
func (s *KnowledgeBaseService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
