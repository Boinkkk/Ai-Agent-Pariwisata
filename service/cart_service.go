package service

import (
	"context"
	"tutorial/models"
	"tutorial/repository"
)

type CartServiceInterface interface {
	Create(ctx context.Context, cart *models.Cart) error
	GetByID(ctx context.Context, id string) (*models.Cart, error)
	GetAll(ctx context.Context) ([]models.Cart, error)
	Update(ctx context.Context, id string, cart *models.Cart) error
	Delete(ctx context.Context, id string) error
}

type CartService struct {
	repo repository.CartRepositoryInterface
}

func NewCartService(repo repository.CartRepositoryInterface) *CartService {
	return &CartService{repo: repo}
}
func (s *CartService) Create(ctx context.Context, cart *models.Cart) error {
	return s.repo.Insert(ctx, cart)
}
func (s *CartService) GetByID(ctx context.Context, id string) (*models.Cart, error) {
	return s.repo.FindByID(ctx, id)
}
func (s *CartService) GetAll(ctx context.Context) ([]models.Cart, error) { return s.repo.FindAll(ctx) }
func (s *CartService) Update(ctx context.Context, id string, cart *models.Cart) error {
	return s.repo.Update(ctx, id, cart)
}
func (s *CartService) Delete(ctx context.Context, id string) error { return s.repo.Delete(ctx, id) }
