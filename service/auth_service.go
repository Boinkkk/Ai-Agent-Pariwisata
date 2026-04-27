package service

import (
	"context"
	"fmt"
	"tutorial/models"
	"tutorial/repository"
	"tutorial/utils"
)

type AuthServiceInterface interface {
	UserLogin(ctx context.Context, email string, password string) (*models.User, error)
}

type AuthService struct {
	repo repository.AuthRepositoryInterface
}

func NewAuthService(repo repository.AuthRepositoryInterface) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) UserLogin(ctx context.Context, email string, password string) (*models.User, error) {
	user, err := s.repo.GetPasswordByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !utils.VerifyPassword(password, user.Password) {
		return nil, fmt.Errorf("wrong password")
	}

	return user, nil
}
