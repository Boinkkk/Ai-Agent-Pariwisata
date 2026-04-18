package service

import (
	"context"
	"tutorial/models"
	"tutorial/repository"
	"tutorial/utils"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashed

	return s.repo.CreateUser(ctx, user)
}
