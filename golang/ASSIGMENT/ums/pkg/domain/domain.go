package domain

import (
	"ums/pkg/db"
	"ums/pkg/exception"
	"ums/pkg/models"
)

// User ...
type User interface {
	GetHealth() bool
	CreateUser(user *models.User) (string, *exception.Exception)
	GetUser(ID, email string) (*models.User, *exception.Exception)
	GetUsers() ([]models.User, *exception.Exception)
	UpdateUser(ID, name string) *exception.Exception
}

// UserClient ...
type UserCliet struct {
	DB      db.Database
	Timeout int
}

// NewUserClient ...
func NewUserCliet(db db.Database, timeout int) *UserCliet {
	return &UserCliet{
		DB:      db,
		Timeout: timeout,
	}
}
