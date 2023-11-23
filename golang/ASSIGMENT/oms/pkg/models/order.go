package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

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

func (p Products) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Make the Products type implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (p *Products) Scan(value interface{}) error {
	var b []byte
	switch t := value.(type) {
	case []byte:
		b = t
	case string:
		b = []byte(t)
	default:
		return errors.New("unknown type")
	}

	return json.Unmarshal(b, &p)
}
