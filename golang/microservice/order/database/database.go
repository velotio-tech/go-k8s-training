package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DataSource struct {
	Mongo *mongo.Client
}

func Init() (*DataSource, error) {

	mongo, err := initMongoDB()
	if err != nil {
		return nil, err
	}

	ds := DataSource{
		Mongo: mongo,
	}

	return &ds, nil

}

func initMongoDB() (*mongo.Client, error) {

	uri := "mongodb://mongodb:27017"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}
