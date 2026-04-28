package service

import (
	"context"
	"tutorial/models"
	"tutorial/repository"
)

type ProductServiceInterface interface {
	Create(ctx context.Context, product *models.Product) error
	GetByID(ctx context.Context, id string) (*models.Product, error)
	GetAll(ctx context.Context) ([]models.Product, error)
	Update(ctx context.Context, id string, product *models.Product) error
	Delete(ctx context.Context, id string) error
	GetBySlug(ctx context.Context, slug string) (*models.Product, error)
}

type ProductService struct {
	repo repository.ProductRepositoryInterface
}

func NewProductService(repo repository.ProductRepositoryInterface) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(ctx context.Context, product *models.Product) error {
	return s.repo.Insert(ctx, product)
}

func (s *ProductService) GetByID(ctx context.Context, id string) (*models.Product, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ProductService) GetBySlug(ctx context.Context, slug string) (*models.Product, error) {
	return s.repo.GetProductBySlug(ctx, slug)
}

func (s *ProductService) GetAll(ctx context.Context) ([]models.Product, error) {
	return s.repo.FindAll(ctx)
}

func (s *ProductService) Update(ctx context.Context, id string, product *models.Product) error {
	return s.repo.Update(ctx, id, product)
}

func (s *ProductService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
