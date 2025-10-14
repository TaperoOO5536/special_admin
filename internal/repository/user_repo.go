package repository

import (
	"context"

	"github.com/TaperoOO5536/special_admin/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers(ctx context.Context, pagination models.Pagination) (*models.PaginatedUsers, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUsers(ctx context.Context, pagination models.Pagination) (*models.PaginatedUsers, error) {
	var users []models.User
	var total int64

	if err := r.db.Model(&models.Item{}).Count(&total).Error; err != nil {
    return nil, err
  }

	offset := (pagination.Page - 1) * pagination.PerPage
	if err := r.db.Limit(pagination.PerPage).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return &models.PaginatedUsers{
		Users:      users,
		TotalCount: total,
		Page:       pagination.Page,
		PerPage:    pagination.PerPage,
	}, nil
}