package service

import (
	"context"
	"errors"

	"github.com/TaperoOO5536/special_admin/internal/models"
	"github.com/TaperoOO5536/special_admin/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrItemPictureNotFound = errors.New("item picture not found")
)

type ItemPictureService struct {
	itemPictureRepo repository.ItemPictureRepository
}

func NewItemPictureService(itemPictureRepo repository.ItemPictureRepository) *ItemPictureService {
	return &ItemPictureService{
		itemPictureRepo: itemPictureRepo,
	}
}

func (s *ItemPictureService) CreateItemPicture(ctx context.Context, itemPicture *models.ItemPicture) (*models.Item, error) {
	item, err := s.itemPictureRepo.CreateItemPicture(ctx, itemPicture)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *ItemPictureService) DeleteItemPicture(ctx context.Context, id uuid.UUID) error {
	err := s.itemPictureRepo.GetItemPicture(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrItemPictureNotFound
		}
		return err
	}

	err = s.itemPictureRepo.DeleteItemPicture(ctx, id)
	if err != nil {
		return err
	}

	return nil
}