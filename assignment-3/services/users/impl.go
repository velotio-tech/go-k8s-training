package users

import (
	model "github.com/pratikpjain/go-k8s-training/assignment3/models"
	repository "github.com/pratikpjain/go-k8s-training/assignment3/repositories"
)

func getAllUsers() ([]model.User, error) {
	users, err := repository.GetAllUsers()
	return users, err
}

func getUserByUserName(username string) (model.User, error) {
	user, err := repository.GetUserByUserName(username)
	return user, err
}

func addNewUser(user model.User) (model.User, error) {
	userCreated, err := repository.AddNewUser(user)
	return userCreated, err
}

func updateUser(user model.User) (model.User, error) {
	oldUser, err := getUserByUserName(user.UserName)
	if err != nil {
		return model.User{}, err
	}
	if user.PhoneNumber == "" {
		user.PhoneNumber = oldUser.PhoneNumber
	}
	if user.City == "" {
		user.City = oldUser.City
	}
	updatedUser, err := repository.UpdateUserByUserName(user)
	if err != nil {
		return model.User{}, err
	}
	return updatedUser, nil
}

func deleteUser(username string) error {
	_, err := getUserByUserName(username)
	if err != nil {
		return err
	}
	err = repository.DeleteUserByUserName(username)
	return err
}
