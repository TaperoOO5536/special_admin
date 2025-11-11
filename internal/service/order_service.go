package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/TaperoOO5536/special_admin/internal/kafka"
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
	userRepo  repository.UserRepository
	producer  *kafka.Producer
}

func NewOrderService(orderRepo repository.OrderRepository,
										 userRepo  repository.UserRepository,
										 producer *kafka.Producer) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
		userRepo:  userRepo,
		producer:  producer,
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

func (s *OrderService) GetOrders(ctx context.Context, pagination models.Pagination) (*models.PaginatedOrders, error) {
	orders, err := s.orderRepo.GetOrders(ctx, pagination)
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

	user, err := s.userRepo.GetUserInfo(ctx, order.UserID)
	if err != nil {
		return nil, err
	}

	go func() {
		msg := models.KafkaOrder{
			Number: order.Number,
			UserID: user.ID,
			Status: order.Status,
		}
		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			log.Printf("failed to marshal message: %v", err)
				return
		}

		err = s.producer.Produce(
			string(jsonMsg),
			"orders",
			"order.update",
		)
		if err != nil {
				log.Printf("failed to produce message: %v", err)
				return
		}
	}()

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

	order, err := s.GetOrderInfo(ctx, id)
	if err != nil {
		return err
	}
	user, err := s.userRepo.GetUserInfo(ctx, order.UserID)
	if err != nil {
		return err
	}

	err = s.orderRepo.DeleteOrder(ctx, id)
	if err != nil {
		return err
	}

	go func() {
		msg := models.KafkaOrder{
			Number: order.Number,
			UserID: user.ID,
		}
		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			log.Printf("failed to marshal message: %v", err)
				return
		}

		err = s.producer.Produce(
			string(jsonMsg),
			"orders",
			"order.delete",
		)
		if err != nil {
				log.Printf("failed to produce message: %v", err)
				return
		}
	}()

	return nil
}