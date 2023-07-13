package model

import "time"

type User struct {
	UserName    string    `json:"username"`
	PhoneNumber string    `json:"phonenumber"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
