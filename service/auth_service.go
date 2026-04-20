package service

import (
	"context"
	"fmt"
	"tutorial/models"
	"tutorial/repository"
	"tutorial/utils"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) UserLogin(ctx context.Context, email string, password string) (*models.User, error) {
	user, err := s.repo.GetPasswordByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if !utils.VerifyPassword(password, user.Password) {
		return nil, fmt.Errorf("Wrong Password")
	}

	return user, nil

}
