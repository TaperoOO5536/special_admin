package repository

import (
	"context"

	"github.com/TaperoOO5536/special_admin/internal/models"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type ItemRepository interface {
	GetItemInfo(ctx context.Context, id uuid.UUID) (*models.Item, error)
	GetItems(ctx context.Context, pagination models.Pagination) (*models.PaginatedItems, error)
	CreateItem(ctx context.Context, item *models.Item) (*models.Item, error)
	UpdateItem(ctx context.Context, item *models.Item, isPriceUpdated bool) (*models.Item, error)
	DeleteItem(ctx context.Context, id uuid.UUID) (error)
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) GetItemInfo(ctx context.Context, id uuid.UUID) (*models.Item, error) {
	var item models.Item
	if err := r.db.
		Preload("Pictures", func(db *gorm.DB) *gorm.DB {
			return db.Select("id_item_picture", "item_id", "picture_path", "mime_type")
		}).
		Where("id_item = ?", id).
		First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *itemRepository) GetItems(ctx context.Context, pagination models.Pagination) (*models.PaginatedItems, error) {
	var items []models.Item
	var total int64

	if err := r.db.Model(&models.Item{}).Count(&total).Error; err != nil {
    return nil, err
  }

	offset := (pagination.Page - 1) * pagination.PerPage
	if err := r.db.Limit(pagination.PerPage).Offset(offset).Find(&items).Error; err != nil {
		return nil, err
	}
	return &models.PaginatedItems{
		Items:      items,
		TotalCount: total,
		Page:       pagination.Page,
		PerPage:    pagination.PerPage,
	}, nil
}

func (r *itemRepository) CreateItem(ctx context.Context, item *models.Item) (*models.Item, error) {
	if err := r.db.Create(item).Error; err != nil {
		return nil, err
	}

	item, err := r.GetItemInfo(ctx, item.ID)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (r *itemRepository) UpdateItem(ctx context.Context, item *models.Item, isPriceUpdated bool) (*models.Item, error) {
	existingItem, err := r.GetItemInfo(ctx, item.ID)
	if err != nil {
		return nil, err
	}

	if item.Title != "" {
		existingItem.Title = item.Title
	}
	if item.Description != "" {
		existingItem.Description = item.Description
	}
	if item.LittlePicture != nil {
		existingItem.LittlePicture = item.LittlePicture
	}
	if isPriceUpdated {
		existingItem.Price = item.Price
	}

	if err := r.db.Save(existingItem).Error; err != nil {
		return nil, err
	}
	item, err = r.GetItemInfo(ctx, item.ID)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (r *itemRepository) DeleteItem(ctx context.Context, id uuid.UUID) error {
	if err := r.db.Where("id_item = ?", id).Delete(&models.Item{}).Error; err != nil {
		return err
	}

	return nil
}