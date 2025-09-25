package repository

import (
	"context"

	"github.com/TaperoOO5536/special_admin/internal/models"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

// var (	ErrNotEnoughSeats = errors.New("ivent does not have enough seats"))

type IventRepository interface {
	GetIventInfo(ctx context.Context, id uuid.UUID) (*models.Ivent, error)
	GetIvents(ctx context.Context) ([]*models.Ivent, error)
	CreateIvent(ctx context.Context, ivent *models.Ivent) (*models.Ivent, error)
	UpdateIvent(ctx context.Context, ivent *models.Ivent) (*models.Ivent, error)
	DeleteIvent(ctx context.Context, id uuid.UUID) (error)
}

type iventRepository struct {
	db *gorm.DB
}

func NewIventRepository(db *gorm.DB) IventRepository {
	return &iventRepository{db: db}
}

func (r *iventRepository) GetIventInfo(ctx context.Context, id uuid.UUID) (*models.Ivent, error) {
	var ivent models.Ivent
	if err := r.db.
		Preload("Pictures", func(db *gorm.DB) *gorm.DB {
			return db.Select("id_ivent_picture", "ivent_id", "picture_path", "mime_type")
		}).
		Preload("UserIvents.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("n_n_user")
		}).
		Where("id_ivent = ?", id).
		First(&ivent).Error; err != nil {
		return nil, err
	}
	return &ivent, nil
}

func (r *iventRepository) GetIvents(ctx context.Context) ([]*models.Ivent, error) {
	var ivents []*models.Ivent
	if err := r.db.Find(&ivents).Error; err != nil {
		return nil, err
	}
	return ivents, nil
}

func (r *iventRepository) CreateIvent(ctx context.Context, ivent *models.Ivent) (*models.Ivent, error) {
	if err := r.db.Create(ivent).Error; err != nil {
		return nil, err
	}
	return ivent, nil
}

func (r *iventRepository) UpdateIvent(ctx context.Context, ivent *models.Ivent) (*models.Ivent, error) {
	existingIvent, err := r.GetIventInfo(ctx, ivent.ID)
	if err != nil {
		return nil, err
	}

	if ivent.Title != "" {
		existingIvent.Title = ivent.Title
	}
	if ivent.Description != "" {
		existingIvent.Description = ivent.Description
	}
	if !ivent.DateTime.IsZero() {
		existingIvent.DateTime = ivent.DateTime
	}
	if ivent.Price != 0 {
		existingIvent.Price = ivent.Price
	}
	if ivent.TotalSeats != 0 {
		existingIvent.TotalSeats = ivent.TotalSeats
	}
	if ivent.OccupiedSeats != 0 {
		existingIvent.OccupiedSeats = ivent.OccupiedSeats
	}
	if ivent.LittlePicture != nil {
		existingIvent.LittlePicture = ivent.LittlePicture
	}

	if err := r.db.Save(existingIvent).Error; err != nil {
		return nil, err
	}
	return ivent, nil
}

func (r *iventRepository) DeleteIvent(ctx context.Context, id uuid.UUID) error {
	if err := r.db.Where("id_ivent = ?", id).Delete(&models.Ivent{}).Error; err != nil {
		return err
	}

	return nil
}