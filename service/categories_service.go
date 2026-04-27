package service

import (
	"context"
	"tutorial/models"
	"tutorial/repository"
)

type CategoriesServiceInterface interface {
	Create(ctx context.Context, category *models.Categories) error
	GetByID(ctx context.Context, id int) (*models.Categories, error)
	GetAll(ctx context.Context) ([]models.Categories, error)
	Update(ctx context.Context, id int, category *models.Categories) error
	Delete(ctx context.Context, id int) error
}

type CategoriesService struct {
	repo repository.CategoriesRepositoryInterface
}

func NewCategoriesService(repo repository.CategoriesRepositoryInterface) *CategoriesService {
	return &CategoriesService{repo: repo}
}

func (s *CategoriesService) Create(ctx context.Context, category *models.Categories) error {
	return s.repo.Insert(ctx, category)
}

func (s *CategoriesService) GetByID(ctx context.Context, id int) (*models.Categories, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *CategoriesService) GetAll(ctx context.Context) ([]models.Categories, error) {
	return s.repo.FindAll(ctx)
}

func (s *CategoriesService) Update(ctx context.Context, id int, category *models.Categories) error {
	return s.repo.Update(ctx, id, category)
}

func (s *CategoriesService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
