package service

import (
	"context"
	"tutorial/models"
	"tutorial/repository"
)

type ProductReviewServiceInterface interface {
	Create(ctx context.Context, review *models.ProductReview) error
	GetByID(ctx context.Context, id string) (*models.ProductReview, error)
	GetAll(ctx context.Context) ([]models.ProductReview, error)
	Update(ctx context.Context, id string, review *models.ProductReview) error
	Delete(ctx context.Context, id string) error
}

type ProductReviewService struct {
	repo repository.ProductReviewRepositoryInterface
}

func NewProductReviewService(repo repository.ProductReviewRepositoryInterface) *ProductReviewService {
	return &ProductReviewService{repo: repo}
}
func (s *ProductReviewService) Create(ctx context.Context, review *models.ProductReview) error {
	return s.repo.Insert(ctx, review)
}
func (s *ProductReviewService) GetByID(ctx context.Context, id string) (*models.ProductReview, error) {
	return s.repo.FindByID(ctx, id)
}
func (s *ProductReviewService) GetAll(ctx context.Context) ([]models.ProductReview, error) {
	return s.repo.FindAll(ctx)
}
func (s *ProductReviewService) Update(ctx context.Context, id string, review *models.ProductReview) error {
	return s.repo.Update(ctx, id, review)
}
func (s *ProductReviewService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
