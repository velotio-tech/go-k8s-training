package models

import "time"

// Order ...
type Order struct {
	ID          string     `json:"id,omitempty"`
	BuyerUserID string     `json:"buyer_user_id,omitempty"`
	Status      string     `json:"status,omitempty"`
	TotalAmount uint64     `json:"total_amount,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	Products    Products   `json:"products,omitempty"`
	PaymentID   string     `json:"payment_id,omitempty"`
}

// Product ...
type Product struct {
	ID       string  `json:"id"`
	Quantity int32   `json:"quantity"`
	Price    float64 `json:"price"`
}

type Products []Product
