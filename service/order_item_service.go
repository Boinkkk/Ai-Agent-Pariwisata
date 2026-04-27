package service

import (
	"context"
	"tutorial/models"
	"tutorial/repository"
)

type OrderItemServiceInterface interface {
	Create(ctx context.Context, item *models.OrderItem) error
	GetByID(ctx context.Context, id string) (*models.OrderItem, error)
	GetAll(ctx context.Context) ([]models.OrderItem, error)
	Update(ctx context.Context, id string, item *models.OrderItem) error
	Delete(ctx context.Context, id string) error
}

type OrderItemService struct {
	repo repository.OrderItemRepositoryInterface
}

func NewOrderItemService(repo repository.OrderItemRepositoryInterface) *OrderItemService {
	return &OrderItemService{repo: repo}
}
func (s *OrderItemService) Create(ctx context.Context, item *models.OrderItem) error {
	return s.repo.Insert(ctx, item)
}
func (s *OrderItemService) GetByID(ctx context.Context, id string) (*models.OrderItem, error) {
	return s.repo.FindByID(ctx, id)
}
func (s *OrderItemService) GetAll(ctx context.Context) ([]models.OrderItem, error) {
	return s.repo.FindAll(ctx)
}
func (s *OrderItemService) Update(ctx context.Context, id string, item *models.OrderItem) error {
	return s.repo.Update(ctx, id, item)
}
func (s *OrderItemService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
