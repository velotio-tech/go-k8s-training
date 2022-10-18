package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/velotio-tech/go-k8s-training/order/model"
	"github.com/velotio-tech/go-k8s-training/order/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type orderService struct {
	orderRepo repository.Order
}

func NewOrderService() Order {
	return &orderService{
		orderRepo: repository.NewOrderRepo(),
	}
}

func (u *orderService) GetAllByUser(ctx context.Context, userID string) ([]model.Order, error) {
	primitiveID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return []model.Order{}, err
	}
	return u.orderRepo.GetAllByUser(ctx, primitiveID)
}

func (u *orderService) CreateOne(ctx context.Context, order model.Order) (model.Order, error) {
	if strings.TrimSpace(order.Product) == "" {
		return model.Order{}, fmt.Errorf("name validation failed")
	}
	return u.orderRepo.CreateOne(ctx, order)
}

func (u *orderService) UpdateOne(ctx context.Context, id string, order model.Order) (model.Order, error) {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Order{}, err
	}
	return u.orderRepo.UpdateOne(ctx, primitiveID, order)
}

func (u *orderService) DeleteOne(ctx context.Context, id string) error {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return u.orderRepo.DeleteOne(ctx, primitiveID)
}

func (u *orderService) DeleteAll(ctx context.Context, id string) error {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return u.orderRepo.DeleteAll(ctx, primitiveID)
}
