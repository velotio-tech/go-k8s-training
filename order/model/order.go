package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	UserID  primitive.ObjectID `json:"user_id" bson:"user_id"`
	Product string             `json:"product" bson:"product"`
}
