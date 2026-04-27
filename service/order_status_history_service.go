package service

import (
	"context"
	"tutorial/models"
	"tutorial/repository"
)

type OrderStatusHistoryServiceInterface interface {
	Create(ctx context.Context, history *models.OrderStatusHistory) error
	GetByID(ctx context.Context, id string) (*models.OrderStatusHistory, error)
	GetAll(ctx context.Context) ([]models.OrderStatusHistory, error)
	Update(ctx context.Context, id string, history *models.OrderStatusHistory) error
	Delete(ctx context.Context, id string) error
}

type OrderStatusHistoryService struct {
	repo repository.OrderStatusHistoryRepositoryInterface
}

func NewOrderStatusHistoryService(repo repository.OrderStatusHistoryRepositoryInterface) *OrderStatusHistoryService {
	return &OrderStatusHistoryService{repo: repo}
}
func (s *OrderStatusHistoryService) Create(ctx context.Context, history *models.OrderStatusHistory) error {
	return s.repo.Insert(ctx, history)
}
func (s *OrderStatusHistoryService) GetByID(ctx context.Context, id string) (*models.OrderStatusHistory, error) {
	return s.repo.FindByID(ctx, id)
}
func (s *OrderStatusHistoryService) GetAll(ctx context.Context) ([]models.OrderStatusHistory, error) {
	return s.repo.FindAll(ctx)
}
func (s *OrderStatusHistoryService) Update(ctx context.Context, id string, history *models.OrderStatusHistory) error {
	return s.repo.Update(ctx, id, history)
}
func (s *OrderStatusHistoryService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
