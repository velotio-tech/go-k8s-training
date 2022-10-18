package repository

import (
	"context"

	"github.com/velotio-tech/go-k8s-training/user/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User interface {
	GetAll(context.Context) ([]model.User, error)
	Create(context.Context, model.User) (model.User, error)
	Update(context.Context, primitive.ObjectID, model.User) (model.User, error)
	Delete(context.Context, primitive.ObjectID) error
}
