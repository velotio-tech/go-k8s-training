package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/velotio-tech/go-k8s-training/user/model"
	"github.com/velotio-tech/go-k8s-training/user/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userService struct {
	userRepo repository.User
}

func NewUserService() User {
	return &userService{
		userRepo: repository.NewUserRepo(),
	}
}

func (u *userService) GetAll(ctx context.Context) ([]model.User, error) {
	return u.userRepo.GetAll(ctx)
}

func (u *userService) Create(ctx context.Context, user model.User) (model.User, error) {
	if strings.TrimSpace(user.Name) == "" {
		return model.User{}, fmt.Errorf("name validation failed")
	}
	return u.userRepo.Create(ctx, user)
}

func (u *userService) Update(ctx context.Context, id string, user model.User) (model.User, error) {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.User{}, err
	}
	return u.userRepo.Update(ctx, primitiveID, user)
}

func (u *userService) Delete(ctx context.Context, id string) error {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return u.userRepo.Delete(ctx, primitiveID)
}
