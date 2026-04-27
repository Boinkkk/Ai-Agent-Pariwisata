package service

import (
	"context"
	"tutorial/models"
	"tutorial/repository"
)

type AddressesServiceInterface interface {
	Create(ctx context.Context, address *models.Addresses) error
	GetByID(ctx context.Context, id string) (*models.Addresses, error)
	GetAll(ctx context.Context) ([]models.Addresses, error)
	Update(ctx context.Context, id string, address *models.Addresses) error
	Delete(ctx context.Context, id string) error
}

type AddressesService struct {
	repo repository.AddressesRepositoryInterface
}

func NewAddressesService(repo repository.AddressesRepositoryInterface) *AddressesService {
	return &AddressesService{repo: repo}
}

func (s *AddressesService) Create(ctx context.Context, address *models.Addresses) error {
	return s.repo.Insert(ctx, address)
}
func (s *AddressesService) GetByID(ctx context.Context, id string) (*models.Addresses, error) {
	return s.repo.FindByID(ctx, id)
}
func (s *AddressesService) GetAll(ctx context.Context) ([]models.Addresses, error) {
	return s.repo.FindAll(ctx)
}
func (s *AddressesService) Update(ctx context.Context, id string, address *models.Addresses) error {
	return s.repo.Update(ctx, id, address)
}
func (s *AddressesService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
