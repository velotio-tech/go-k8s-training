package repository

import (
	"context"

	"github.com/velotio-tech/go-k8s-training/user/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepo struct{}

func NewUserRepo() User {
	return &userRepo{}
}

func (u *userRepo) GetAll(ctx context.Context) ([]model.User, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database("users").Collection("records")
	filter := bson.M{}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	users := []model.User{}
	err = cursor.All(ctx, &users)
	return users, err
}

func (u *userRepo) Create(ctx context.Context, user model.User) (model.User, error) {
	client, err := getClient()
	if err != nil {
		return model.User{}, err
	}
	user.ID = primitive.NewObjectID()
	collection := client.Database("users").Collection("records")
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *userRepo) Update(ctx context.Context, id primitive.ObjectID, user model.User) (model.User, error) {
	client, err := getClient()
	if err != nil {
		return model.User{}, err
	}
	collection := client.Database("users").Collection("records")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"name": user.Name}}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	updatedUser := model.User{}
	err = collection.FindOneAndUpdate(ctx, filter, update, options).Decode(&u)
	return updatedUser, err
}

func (u *userRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
	client, err := getClient()
	if err != nil {
		return err
	}
	collection := client.Database("users").Collection("records")
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
