package repository

import (
	"context"

	"github.com/TaperoOO5536/special_admin/internal/models"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type ItemPictureRepository interface {
	CreateItemPicture(ctx context.Context, itemPicture *models.ItemPicture) (*models.Item, error)
	DeleteItemPicture(ctx context.Context, id uuid.UUID) (error)
  GetItemPicture(ctx context.Context, id uuid.UUID) error
}

type itemPictureRepository struct {
	db *gorm.DB
	itemRepo ItemRepository
}

func NewItemPictureRepository(db *gorm.DB, itemRepo ItemRepository) ItemPictureRepository {
	return &itemPictureRepository{
		db: db,
	  itemRepo: itemRepo,
	}
}

func (r *itemPictureRepository) CreateItemPicture(ctx context.Context, itemPicture *models.ItemPicture) (*models.Item, error) {
	if err := r.db.Create(itemPicture).Error; err != nil {
		return nil, err
	}

	item, err := r.itemRepo.GetItemInfo(ctx, itemPicture.ItemID)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (r *itemPictureRepository) DeleteItemPicture(ctx context.Context, id uuid.UUID) (error) {
	if err := r.db.Where("id_item_picture = ?", id).Delete(&models.ItemPicture{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *itemPictureRepository) GetItemPicture(ctx context.Context, id uuid.UUID) error {
	var itemPicture models.ItemPicture
	if err := r.db.Where("id_item_picture = ?", id).First(&itemPicture).Error; err != nil {
		return err
	}
	return nil
}