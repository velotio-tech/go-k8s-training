package service

import (
	"context"

	"github.com/velotio-tech/go-k8s-training/user/model"
)

type User interface {
	GetAll(context.Context) ([]model.User, error)
	Create(context.Context, model.User) (model.User, error)
	Update(context.Context, string, model.User) (model.User, error)
	Delete(context.Context, string) error
}
