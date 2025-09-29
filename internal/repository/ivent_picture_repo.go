package repository

import (
	"context"

	"github.com/TaperoOO5536/special_admin/internal/models"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type IventPictureRepository interface {
	CreateIventPicture(ctx context.Context, iventPicture *models.IventPicture) (*models.Ivent, error)
	DeleteIventPicture(ctx context.Context, id uuid.UUID) (error)
  GetIventPicture(ctx context.Context, id uuid.UUID) error
}

type iventPictureRepository struct {
	db *gorm.DB
	iventRepo IventRepository
}

func NewIventPictureRepository(db *gorm.DB, iventRepo IventRepository) IventPictureRepository {
	return &iventPictureRepository{
		db: db,
	  iventRepo: iventRepo,
	}
}

func (r *iventPictureRepository) CreateIventPicture(ctx context.Context, iventPicture *models.IventPicture) (*models.Ivent, error) {
	if err := r.db.Create(iventPicture).Error; err != nil {
		return nil, err
	}

	ivent, err := r.iventRepo.GetIventInfo(ctx, iventPicture.IventID)
	if err != nil {
		return nil, err
	}
	return ivent, nil
}

func (r *iventPictureRepository) DeleteIventPicture(ctx context.Context, id uuid.UUID) (error) {
	if err := r.db.Where("id_ivent_picture = ?", id).Delete(&models.IventPicture{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *iventPictureRepository) GetIventPicture(ctx context.Context, id uuid.UUID) error {
	var iventPicture models.IventPicture
	if err := r.db.Where("id_ivent_picture = ?", id).First(&iventPicture).Error; err != nil {
		return err
	}
	return nil
}