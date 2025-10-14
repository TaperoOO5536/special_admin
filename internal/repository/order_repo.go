package repository

import (
	"context"

	"github.com/TaperoOO5536/special_admin/internal/models"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type OrderRepository interface {
	GetOrderInfo(ctx context.Context, id uuid.UUID) (*models.Order, error)
	GetOrders(ctx context.Context, pagination models.Pagination) (*models.PaginatedOrders, error)
	UpdateOrder(ctx context.Context, order *models.Order) (*models.Order, error)
	DeleteOrder(ctx context.Context, id uuid.UUID) (error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) GetOrderInfo(ctx context.Context, id uuid.UUID) (*models.Order, error) {
	var order models.Order
	if err := r.db.
		Preload("OrderItems.Item", func(db *gorm.DB) *gorm.DB {
			return db.Select("id_item", "item_title", "item_price", "little_picture", "mime_type")
		}).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id_user", "n_n_user")
		}).
		Where("id_order = ?", id).
		First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) GetOrders(ctx context.Context, pagination models.Pagination) (*models.PaginatedOrders, error) {
	var orders []models.Order
	var total int64
	
	if err := r.db.Model(&models.Order{}).Count(&total).Error; err != nil {
    return nil, err
  }

	offset := (pagination.Page - 1) * pagination.PerPage
	if err := r.db.Limit(pagination.PerPage).Offset(offset).Find(&orders).Error; err != nil {
		return nil, err
	}
	return &models.PaginatedOrders{
		Orders:     orders,
		TotalCount: total,
		Page:       pagination.Page,
		PerPage:    pagination.PerPage,
	}, nil
}

func (r *orderRepository) UpdateOrder(ctx context.Context, order *models.Order) (*models.Order, error) {
	existingOrder, err := r.GetOrderInfo(ctx, order.ID)
	if err != nil {
		return nil, err
	}
	
	if order.Status != "" {
		existingOrder.Status = order.Status
	}
	if order.Comment != "" {
		existingOrder.Comment = order.Comment
	}	

	if err := r.db.Save(existingOrder).Error; err != nil {
		return nil, err
	}
	order, err = r.GetOrderInfo(ctx, order.ID)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *orderRepository) DeleteOrder(ctx context.Context, id uuid.UUID) error {
	if err := r.db.Where("id_order = ?", id).Delete(&models.Order{}).Error; err != nil {
		return err
	}

	return nil
}