package service

import (
	"context"
	"tutorial/models"
	"tutorial/repository"
)

type RoleServiceInterface interface {
	Create(ctx context.Context, role *models.Role) error
	GetByID(ctx context.Context, id int) (*models.Role, error)
	GetAll(ctx context.Context) ([]models.Role, error)
	Update(ctx context.Context, id int, role *models.Role) error
	Delete(ctx context.Context, id int) error
}

type RoleService struct {
	repo repository.RoleRepositoryInterface
}

func NewRoleService(repo repository.RoleRepositoryInterface) *RoleService {
	return &RoleService{repo: repo}
}

func (s *RoleService) Create(ctx context.Context, role *models.Role) error {
	return s.repo.Insert(ctx, role)
}

func (s *RoleService) GetByID(ctx context.Context, id int) (*models.Role, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *RoleService) GetAll(ctx context.Context) ([]models.Role, error) {
	return s.repo.FindAll(ctx)
}

func (s *RoleService) Update(ctx context.Context, id int, role *models.Role) error {
	return s.repo.Update(ctx, id, role)
}

func (s *RoleService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
