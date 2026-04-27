package service

import (
	"context"
	"tutorial/models"
	"tutorial/repository"
)

type PaymentServiceInterface interface {
	Create(ctx context.Context, payment *models.Payment) error
	GetByID(ctx context.Context, id string) (*models.Payment, error)
	GetAll(ctx context.Context) ([]models.Payment, error)
	Update(ctx context.Context, id string, payment *models.Payment) error
	Delete(ctx context.Context, id string) error
}

type PaymentService struct {
	repo repository.PaymentRepositoryInterface
}

func NewPaymentService(repo repository.PaymentRepositoryInterface) *PaymentService {
	return &PaymentService{repo: repo}
}
func (s *PaymentService) Create(ctx context.Context, payment *models.Payment) error {
	return s.repo.Insert(ctx, payment)
}
func (s *PaymentService) GetByID(ctx context.Context, id string) (*models.Payment, error) {
	return s.repo.FindByID(ctx, id)
}
func (s *PaymentService) GetAll(ctx context.Context) ([]models.Payment, error) {
	return s.repo.FindAll(ctx)
}
func (s *PaymentService) Update(ctx context.Context, id string, payment *models.Payment) error {
	return s.repo.Update(ctx, id, payment)
}
func (s *PaymentService) Delete(ctx context.Context, id string) error { return s.repo.Delete(ctx, id) }
