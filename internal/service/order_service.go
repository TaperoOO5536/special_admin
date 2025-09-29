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
	ErrOrderNotFound = errors.New("order not found")
)

type OrderService struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (s *OrderService) GetOrderInfo(ctx context.Context, id uuid.UUID) (*models.Order, error) {

	order, err := s.orderRepo.GetOrderInfo(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}
	
	return order, nil 
}

func (s *OrderService) GetOrders(ctx context.Context) ([]*models.Order, error) {
	orders, err := s.orderRepo.GetOrders(ctx)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *OrderService) UpdateOrder(ctx context.Context, order *models.Order) (*models.Order, error) {

	_, err := s.orderRepo.GetOrderInfo(ctx, order.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}

	order, err = s.orderRepo.UpdateOrder(ctx, order)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}

	return order, nil
}

func (s *OrderService) DeleteOrder(ctx context.Context, id uuid.UUID) error {
	_, err := s.orderRepo.GetOrderInfo(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrOrderNotFound
		}
		return err
	}

	err = s.orderRepo.DeleteOrder(ctx, id)
	if err != nil {
		return err
	}

	return nil
}