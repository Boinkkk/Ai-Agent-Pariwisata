package service

import (
	"context"
	"tutorial/models"
	"tutorial/repository"
)

type ArticleServiceInterface interface {
	Create(ctx context.Context, article *models.Article) error
	GetByID(ctx context.Context, id string) (*models.Article, error)
	GetAll(ctx context.Context) ([]models.Article, error)
	Update(ctx context.Context, id string, article *models.Article) error
	Delete(ctx context.Context, id string) error
}

type ArticleService struct {
	repo repository.ArticleRepositoryInterface
}

func NewArticleService(repo repository.ArticleRepositoryInterface) *ArticleService {
	return &ArticleService{repo: repo}
}
func (s *ArticleService) Create(ctx context.Context, article *models.Article) error {
	return s.repo.Insert(ctx, article)
}
func (s *ArticleService) GetByID(ctx context.Context, id string) (*models.Article, error) {
	return s.repo.FindByID(ctx, id)
}
func (s *ArticleService) GetAll(ctx context.Context) ([]models.Article, error) {
	return s.repo.FindAll(ctx)
}
func (s *ArticleService) Update(ctx context.Context, id string, article *models.Article) error {
	return s.repo.Update(ctx, id, article)
}
func (s *ArticleService) Delete(ctx context.Context, id string) error { return s.repo.Delete(ctx, id) }
