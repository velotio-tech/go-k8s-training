package repository

import (
	"context"

	"github.com/velotio-tech/go-k8s-training/order/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order interface {
	GetAllByUser(context.Context, primitive.ObjectID) ([]model.Order, error)
	CreateOne(context.Context, model.Order) (model.Order, error)
	UpdateOne(context.Context, primitive.ObjectID, model.Order) (model.Order, error)
	DeleteOne(context.Context, primitive.ObjectID) error
	DeleteAll(context.Context, primitive.ObjectID) error
}
