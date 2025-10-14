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
	ErrEventNotFound = errors.New("event not found")
)

type EventService struct {
	eventRepo repository.EventRepository
}

func NewEventService(eventRepo repository.EventRepository) *EventService {
	return &EventService{
		eventRepo: eventRepo,
	}
}

func (s *EventService) CreateEvent(ctx context.Context, event *models.Event) (*models.Event, error) {
	event, err := s.eventRepo.CreateEvent(ctx, event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (s *EventService) GetEventInfo(ctx context.Context, id uuid.UUID) (*models.Event, error) {

	event, err := s.eventRepo.GetEventInfo(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrEventNotFound
		}
		return nil, err
	}
	
	return event, nil 
}

func (s *EventService) GetEvents(ctx context.Context, pagination models.Pagination) (*models.PaginatedEvents, error) {
	events, err := s.eventRepo.GetEvents(ctx, pagination)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (s *EventService) UpdateEvent(ctx context.Context, event *models.Event, isPriceUpdated bool, isOccupiedSeats bool) (*models.Event, error) {

	_, err := s.eventRepo.GetEventInfo(ctx, event.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrEventNotFound
		}
		return nil, err
	}

	event, err = s.eventRepo.UpdateEvent(ctx, event, isPriceUpdated, isOccupiedSeats)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrEventNotFound
		}
		return nil, err
	}

	return event, nil
}

func (s *EventService) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	_, err := s.eventRepo.GetEventInfo(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrEventNotFound
		}
		return err
	}

	err = s.eventRepo.DeleteEvent(ctx, id)
	if err != nil {
		return err
	}

	return nil
}