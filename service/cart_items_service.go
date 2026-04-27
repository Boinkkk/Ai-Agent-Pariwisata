package service

import (
	"context"
	"tutorial/models"
	"tutorial/repository"
)

type CartItemServiceInterface interface {
	Create(ctx context.Context, item *models.CartItem) error
	GetByID(ctx context.Context, id string) (*models.CartItem, error)
	GetAll(ctx context.Context) ([]models.CartItem, error)
	Update(ctx context.Context, id string, item *models.CartItem) error
	Delete(ctx context.Context, id string) error
}

type CartItemService struct {
	repo repository.CartItemRepositoryInterface
}

func NewCartItemService(repo repository.CartItemRepositoryInterface) *CartItemService {
	return &CartItemService{repo: repo}
}
func (s *CartItemService) Create(ctx context.Context, item *models.CartItem) error {
	return s.repo.Insert(ctx, item)
}
func (s *CartItemService) GetByID(ctx context.Context, id string) (*models.CartItem, error) {
	return s.repo.FindByID(ctx, id)
}
func (s *CartItemService) GetAll(ctx context.Context) ([]models.CartItem, error) {
	return s.repo.FindAll(ctx)
}
func (s *CartItemService) Update(ctx context.Context, id string, item *models.CartItem) error {
	return s.repo.Update(ctx, id, item)
}
func (s *CartItemService) Delete(ctx context.Context, id string) error { return s.repo.Delete(ctx, id) }
