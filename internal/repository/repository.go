package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	User      UserRepository
	Ivent     IventRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:      NewUserRepository(db),
		Ivent:     NewIventRepository(db),
	}
}