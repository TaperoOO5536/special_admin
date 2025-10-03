package repository

import (
	"context"

	"github.com/TaperoOO5536/special_admin/internal/models"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type EventPictureRepository interface {
	CreateEventPicture(ctx context.Context, eventPicture *models.EventPicture) (*models.Event, error)
	DeleteEventPicture(ctx context.Context, id uuid.UUID) (error)
  GetEventPicture(ctx context.Context, id uuid.UUID) error
}

type eventPictureRepository struct {
	db *gorm.DB
	eventRepo EventRepository
}

func NewEventPictureRepository(db *gorm.DB, eventRepo EventRepository) EventPictureRepository {
	return &eventPictureRepository{
		db: db,
	  eventRepo: eventRepo,
	}
}

func (r *eventPictureRepository) CreateEventPicture(ctx context.Context, eventPicture *models.EventPicture) (*models.Event, error) {
	if err := r.db.Create(eventPicture).Error; err != nil {
		return nil, err
	}

	event, err := r.eventRepo.GetEventInfo(ctx, eventPicture.EventID)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (r *eventPictureRepository) DeleteEventPicture(ctx context.Context, id uuid.UUID) (error) {
	if err := r.db.Where("id_event_picture = ?", id).Delete(&models.EventPicture{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *eventPictureRepository) GetEventPicture(ctx context.Context, id uuid.UUID) error {
	var eventPicture models.EventPicture
	if err := r.db.Where("id_event_picture = ?", id).First(&eventPicture).Error; err != nil {
		return err
	}
	return nil
}