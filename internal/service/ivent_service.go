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
	ErrIventNotFound = errors.New("ivent not found")
)

type IventService struct {
	iventRepo repository.IventRepository
}

func NewIventService(iventRepo repository.IventRepository) *IventService {
	return &IventService{
		iventRepo: iventRepo,
	}
}

func (s *IventService) CreateIvent(ctx context.Context, ivent *models.Ivent) (*models.Ivent, error) {
	ivent, err := s.iventRepo.CreateIvent(ctx, ivent)
	if err != nil {
		return nil, err
	}

	return ivent, nil
}

func (s *IventService) GetIventInfo(ctx context.Context, id uuid.UUID) (*models.Ivent, error) {

	ivent, err := s.iventRepo.GetIventInfo(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrIventNotFound
		}
		return nil, err
	}
	
	return ivent, nil 
}

func (s *IventService) GetIvents(ctx context.Context) ([]*models.Ivent, error) {
	ivents, err := s.iventRepo.GetIvents(ctx)
	if err != nil {
		return nil, err
	}

	return ivents, nil
}

func (s *IventService) UpdateIvent(ctx context.Context, ivent *models.Ivent, isPriceUpdated bool, isOccupiedSeats bool) (*models.Ivent, error) {

	_, err := s.iventRepo.GetIventInfo(ctx, ivent.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrIventNotFound
		}
		return nil, err
	}

	ivent, err = s.iventRepo.UpdateIvent(ctx, ivent, isPriceUpdated, isOccupiedSeats)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrIventNotFound
		}
		return nil, err
	}

	return ivent, nil
}

func (s *IventService) DeleteIvent(ctx context.Context, id uuid.UUID) error {
	_, err := s.iventRepo.GetIventInfo(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrIventNotFound
		}
		return err
	}

	err = s.iventRepo.DeleteIvent(ctx, id)
	if err != nil {
		return err
	}

	return nil
}