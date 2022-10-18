package repository

import (
	"context"

	"github.com/velotio-tech/go-k8s-training/order/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoConnection *mongo.Client

func getClient() (*mongo.Client, error) {
	dbURL := utils.GetServiceConfig().DatabaseURL
	if mongoConnection == nil {
		mongoConnection, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURL))
		return mongoConnection, err
	}
	return mongoConnection, nil
}
