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
	ErrIventPictureNotFound = errors.New("ivent picture not found")
)

type IventPictureService struct {
	iventPictureRepo repository.IventPictureRepository
}

func NewIventPictureService(iventPictureRepo repository.IventPictureRepository) *IventPictureService {
	return &IventPictureService{
		iventPictureRepo: iventPictureRepo,
	}
}

func (s *IventPictureService) CreateIventPicture(ctx context.Context, iventPicture *models.IventPicture) (*models.Ivent, error) {
	ivent, err := s.iventPictureRepo.CreateIventPicture(ctx, iventPicture)
	if err != nil {
		return nil, err
	}

	return ivent, nil
}

func (s *IventPictureService) DeleteIventPicture(ctx context.Context, id uuid.UUID) error {
	err := s.iventPictureRepo.GetIventPicture(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrIventPictureNotFound
		}
		return err
	}

	err = s.iventPictureRepo.DeleteIventPicture(ctx, id)
	if err != nil {
		return err
	}

	return nil
}