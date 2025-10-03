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
	ErrEventPictureNotFound = errors.New("event picture not found")
)

type EventPictureService struct {
	eventPictureRepo repository.EventPictureRepository
}

func NewEventPictureService(eventPictureRepo repository.EventPictureRepository) *EventPictureService {
	return &EventPictureService{
		eventPictureRepo: eventPictureRepo,
	}
}

func (s *EventPictureService) CreateEventPicture(ctx context.Context, eventPicture *models.EventPicture) (*models.Event, error) {
	event, err := s.eventPictureRepo.CreateEventPicture(ctx, eventPicture)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (s *EventPictureService) DeleteEventPicture(ctx context.Context, id uuid.UUID) error {
	err := s.eventPictureRepo.GetEventPicture(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrEventPictureNotFound
		}
		return err
	}

	err = s.eventPictureRepo.DeleteEventPicture(ctx, id)
	if err != nil {
		return err
	}

	return nil
}