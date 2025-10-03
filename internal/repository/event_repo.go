package repository

import (
	"context"

	"github.com/TaperoOO5536/special_admin/internal/models"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type EventRepository interface {
	GetEventInfo(ctx context.Context, id uuid.UUID) (*models.Event, error)
	GetEvents(ctx context.Context) ([]*models.Event, error)
	CreateEvent(ctx context.Context, event *models.Event) (*models.Event, error)
	UpdateEvent(ctx context.Context, event *models.Event, isPriceUpdated bool, isOccupiedSeats bool) (*models.Event, error)
	DeleteEvent(ctx context.Context, id uuid.UUID) (error)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) GetEventInfo(ctx context.Context, id uuid.UUID) (*models.Event, error) {
	var event models.Event
	if err := r.db.
		Preload("Pictures", func(db *gorm.DB) *gorm.DB {
			return db.Select("id_event_picture", "event_id", "picture_path", "mime_type")
		}).
		Preload("UserEvents.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("n_n_user")
		}).
		Where("id_event = ?", id).
		First(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) GetEvents(ctx context.Context) ([]*models.Event, error) {
	var events []*models.Event
	if err := r.db.Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *eventRepository) CreateEvent(ctx context.Context, event *models.Event) (*models.Event, error) {
	if err := r.db.Create(event).Error; err != nil {
		return nil, err
	}

	event, err := r.GetEventInfo(ctx, event.ID)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (r *eventRepository) UpdateEvent(ctx context.Context, event *models.Event, isPriceUpdated bool, isOccupiedSeats bool) (*models.Event, error) {
	existingEvent, err := r.GetEventInfo(ctx, event.ID)
	if err != nil {
		return nil, err
	}

	if event.Title != "" {
		existingEvent.Title = event.Title
	}
	if event.Description != "" {
		existingEvent.Description = event.Description
	}
	if !event.DateTime.IsZero() {
		existingEvent.DateTime = event.DateTime
	}
	if event.TotalSeats != 0 {
		existingEvent.TotalSeats = event.TotalSeats
	}
	if event.OccupiedSeats != 0 {
		existingEvent.OccupiedSeats = event.OccupiedSeats
	}
	if event.LittlePicture != nil {
		existingEvent.LittlePicture = event.LittlePicture
	}
	if isPriceUpdated {
		existingEvent.Price = event.Price
	}
	if isOccupiedSeats {
		existingEvent.OccupiedSeats = event.OccupiedSeats
	}

	if err := r.db.Save(existingEvent).Error; err != nil {
		return nil, err
	}
	event, err = r.GetEventInfo(ctx, event.ID)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (r *eventRepository) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	if err := r.db.Where("id_event = ?", id).Delete(&models.Event{}).Error; err != nil {
		return err
	}

	return nil
}