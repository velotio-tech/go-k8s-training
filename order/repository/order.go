package repository

import (
	"context"

	"github.com/velotio-tech/go-k8s-training/order/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type orderRepo struct{}

func NewOrderRepo() Order {
	return &orderRepo{}
}

func (u *orderRepo) GetAllByUser(ctx context.Context, userID primitive.ObjectID) ([]model.Order, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database("orders").Collection("records")
	filter := bson.M{"user_id": userID}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	orders := []model.Order{}
	err = cursor.All(ctx, &orders)
	return orders, err
}

func (u *orderRepo) CreateOne(ctx context.Context, order model.Order) (model.Order, error) {
	client, err := getClient()
	if err != nil {
		return model.Order{}, err
	}
	order.ID = primitive.NewObjectID()
	collection := client.Database("orders").Collection("records")
	_, err = collection.InsertOne(ctx, order)
	if err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func (u *orderRepo) UpdateOne(ctx context.Context, orderID primitive.ObjectID, order model.Order) (model.Order, error) {
	client, err := getClient()
	if err != nil {
		return model.Order{}, err
	}
	collection := client.Database("orders").Collection("records")
	filter := bson.M{"_id": orderID}
	update := bson.M{"$set": bson.M{"product": order.Product}}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	updatedOrder := model.Order{}
	err = collection.FindOneAndUpdate(ctx, filter, update, options).Decode(&updatedOrder)
	return updatedOrder, err
}

func (u *orderRepo) DeleteOne(ctx context.Context, orderID primitive.ObjectID) error {
	client, err := getClient()
	if err != nil {
		return err
	}
	collection := client.Database("orders").Collection("records")
	filter := bson.M{"_id": orderID}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (u *orderRepo) DeleteAll(ctx context.Context, userID primitive.ObjectID) error {
	client, err := getClient()
	if err != nil {
		return err
	}
	collection := client.Database("orders").Collection("records")
	filter := bson.M{"user_id": userID}
	result, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
