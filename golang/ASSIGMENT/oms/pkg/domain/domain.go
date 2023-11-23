package domain

import (
	"oms/pkg/models"

	"oms/pkg/db"
)

// Order ...
type Order interface {
	GetHealth() bool
	CreateOrder(order *models.Order) (string, error)
	GetOrdersByUserID(userID string) ([]models.Order, error)
	DeleteOrderByUserID(userID string) error
	DeleteOrderByOrderID(orderID string) error
	GetOrder(orderID string) (*models.Order, error)
}

// UserClient ...
type OrderCliet struct {
	DB      db.Database
	Timeout int
}

// NewUserClient ...
func NewOrderCliet(db db.Database, timeout int) *OrderCliet {
	return &OrderCliet{
		DB:      db,
		Timeout: timeout,
	}
}
