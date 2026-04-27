package service

import (
	"context"
	"tutorial/models"
	"tutorial/repository"
)

type OrderServiceInterface interface {
	Create(ctx context.Context, order *models.Order) error
	GetByID(ctx context.Context, id string) (*models.Order, error)
	GetAll(ctx context.Context) ([]models.Order, error)
	Update(ctx context.Context, id string, order *models.Order) error
	Delete(ctx context.Context, id string) error
}

type OrderService struct {
	repo repository.OrderRepositoryInterface
}

func NewOrderService(repo repository.OrderRepositoryInterface) *OrderService {
	return &OrderService{repo: repo}
}
func (s *OrderService) Create(ctx context.Context, order *models.Order) error {
	return s.repo.Insert(ctx, order)
}
func (s *OrderService) GetByID(ctx context.Context, id string) (*models.Order, error) {
	return s.repo.FindByID(ctx, id)
}
func (s *OrderService) GetAll(ctx context.Context) ([]models.Order, error) {
	return s.repo.FindAll(ctx)
}
func (s *OrderService) Update(ctx context.Context, id string, order *models.Order) error {
	return s.repo.Update(ctx, id, order)
}
func (s *OrderService) Delete(ctx context.Context, id string) error { return s.repo.Delete(ctx, id) }
