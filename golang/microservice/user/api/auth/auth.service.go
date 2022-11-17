package auth

import (
	"context"

	"github.com/segmentio/ksuid"
	"github.com/velotio-ajaykumbhar/microservice/user/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Database       = "microservice"
	AuthCollection = "auth"
)

type AuthService interface {
	Register(context.Context, AuthDTO) (Auth, error)
	Login(context.Context, AuthDTO) (Auth, error)
	ReadAll(context.Context) ([]Auth, error)
}

type authService struct {
	Mongo *mongo.Client
}

func NewAuthService(ds *database.DataSource) AuthService {
	return &authService{
		Mongo: ds.Mongo,
	}
}

func (as authService) Register(ctx context.Context, authDTO AuthDTO) (Auth, error) {

	auth := Auth{
		AuthId:   ksuid.New().String(),
		Username: authDTO.Username,
		Password: authDTO.Password,
	}

	authCollection := as.Mongo.Database(Database).Collection(AuthCollection)

	_, err := authCollection.InsertOne(ctx, auth)

	if err != nil {
		return Auth{}, err
	}
	return auth, nil
}

func (as authService) Login(ctx context.Context, authDTO AuthDTO) (Auth, error) {

	authCollection := as.Mongo.Database(Database).Collection(AuthCollection)

	auth := Auth{}

	filter := bson.M{
		"username": authDTO.Username,
		"password": authDTO.Password,
	}
	err := authCollection.FindOne(ctx, filter).Decode(&auth)

	if err != nil {
		return Auth{}, err
	}

	return auth, nil
}

func (as authService) ReadAll(ctx context.Context) ([]Auth, error) {

	authCollection := as.Mongo.Database(Database).Collection(AuthCollection)

	filter := bson.M{}
	cursor, err := authCollection.Find(ctx, filter)
	if err != nil {
		return []Auth{}, err
	}

	auths := []Auth{}
	err = cursor.All(ctx, &auths)
	if err != nil {
		return []Auth{}, err
	}

	return auths, nil
}
