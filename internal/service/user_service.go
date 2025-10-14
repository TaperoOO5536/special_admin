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

func (s *UserService) GetUsers(ctx context.Context, pagination models.Pagination) (*models.PaginatedUsers, error) {
	users, err := s.userRepo.GetUsers(ctx, pagination)
	if err != nil {
		return nil, err
	}
	
	return users, nil
}