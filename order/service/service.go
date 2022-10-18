package service

import (
	"context"

	"github.com/velotio-tech/go-k8s-training/order/model"
)

type Order interface {
	GetAllByUser(context.Context, string) ([]model.Order, error)
	CreateOne(context.Context, model.Order) (model.Order, error)
	UpdateOne(context.Context, string, model.Order) (model.Order, error)
	DeleteOne(context.Context, string) error
	DeleteAll(context.Context, string) error
}
