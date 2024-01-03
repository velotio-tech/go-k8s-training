package model

import "time"

type Order struct {
	OrderID   int       `json:"order_id"`
	OrderDesc string    `json:"order_desc"`
	UserName  string    `json:"username"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
