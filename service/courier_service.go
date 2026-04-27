package service

import (
	"context"
	"tutorial/models"
	"tutorial/repository"
)

type CourierServiceInterface interface {
	Create(ctx context.Context, courier *models.Courier) error
	GetByID(ctx context.Context, id int) (*models.Courier, error)
	GetAll(ctx context.Context) ([]models.Courier, error)
	Update(ctx context.Context, id int, courier *models.Courier) error
	Delete(ctx context.Context, id int) error
}

type CourierService struct {
	repo repository.CourierRepositoryInterface
}

func NewCourierService(repo repository.CourierRepositoryInterface) *CourierService {
	return &CourierService{repo: repo}
}
func (s *CourierService) Create(ctx context.Context, courier *models.Courier) error {
	return s.repo.Insert(ctx, courier)
}
func (s *CourierService) GetByID(ctx context.Context, id int) (*models.Courier, error) {
	return s.repo.FindByID(ctx, id)
}
func (s *CourierService) GetAll(ctx context.Context) ([]models.Courier, error) {
	return s.repo.FindAll(ctx)
}
func (s *CourierService) Update(ctx context.Context, id int, courier *models.Courier) error {
	return s.repo.Update(ctx, id, courier)
}
func (s *CourierService) Delete(ctx context.Context, id int) error { return s.repo.Delete(ctx, id) }
