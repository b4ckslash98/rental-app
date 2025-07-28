package repository

import (
	"github.com/b4ckslash98/rental-app/services/order/entity"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Save(order *entity.Order) error
	FindByUser(userID int) ([]*entity.Order, error)
	FindAll() ([]*entity.Order, error)
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepo{db}
}

func (r *orderRepo) Save(order *entity.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepo) FindByUser(userID int) ([]*entity.Order, error) {
	var orders []*entity.Order
	err := r.db.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (r *orderRepo) FindAll() ([]*entity.Order, error) {
	var orders []*entity.Order
	err := r.db.Find(&orders).Error
	return orders, err
}
