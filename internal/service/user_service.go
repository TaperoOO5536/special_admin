package service

import (
	"context"

	"github.com/TaperoOO5536/special_admin/internal/models"
	"github.com/TaperoOO5536/special_admin/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUsers(ctx context.Context) ([]*models.User, error) {
	users, err := s.userRepo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	
	return users, nil
}