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
	ErrItemNotFound = errors.New("item not found")
)

type ItemService struct {
	itemRepo repository.ItemRepository
}

func NewItemService(itemRepo repository.ItemRepository) *ItemService {
	return &ItemService{
		itemRepo: itemRepo,
	}
}

func (s *ItemService) CreateItem(ctx context.Context, item *models.Item) (*models.Item, error) {
	item, err := s.itemRepo.CreateItem(ctx, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *ItemService) GetItemInfo(ctx context.Context, id uuid.UUID) (*models.Item, error) {

	item, err := s.itemRepo.GetItemInfo(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrItemNotFound
		}
		return nil, err
	}
	
	return item, nil 
}

func (s *ItemService) GetItems(ctx context.Context) ([]*models.Item, error) {
	items, err := s.itemRepo.GetItems(ctx)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *ItemService) UpdateItem(ctx context.Context, item *models.Item, isPriceUpdated bool) (*models.Item, error) {

	_, err := s.itemRepo.GetItemInfo(ctx, item.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrItemNotFound
		}
		return nil, err
	}

	item, err = s.itemRepo.UpdateItem(ctx, item, isPriceUpdated)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrItemNotFound
		}
		return nil, err
	}

	return item, nil
}

func (s *ItemService) DeleteItem(ctx context.Context, id uuid.UUID) error {
	_, err := s.itemRepo.GetItemInfo(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrItemNotFound
		}
		return err
	}

	err = s.itemRepo.DeleteItem(ctx, id)
	if err != nil {
		return err
	}

	return nil
}