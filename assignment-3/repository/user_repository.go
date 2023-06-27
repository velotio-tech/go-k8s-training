package repository

import (
	"time"

	"github.com/practice/db"
	"github.com/practice/model"
)

const (
	getUserByUserName    = `SELECT phonenumber, city, created_at, updated_at FROM USERS WHERE username=$1`
	getAllUsers          = `SELECT username, phonenumber, city, created_at, updated_at FROM USERS`
	insertUser           = `INSERT INTO USERS (username, phonenumber, city, created_at, updated_at) VALUES ($1, $2, $3, now(), now()) RETURNING username, phonenumber, city, created_at, updated_at`
	updateUserByUserName = `UPDATE USERS SET phonenumber = $2, city = $3, updated_at = now() where username = $1 RETURNING username, phonenumber, city, created_at, updated_at`
	deleteUserByUserName = `DELETE FROM USERS WHERE username = $1`
)

func DeleteUserByUserName(username string) error {
	_, err := db.GetDB().Exec(deleteUserByUserName, username)
	return err
}

func UpdateUserByUserName(user model.User) (model.User, error) {

	var updatedUser model.User

	err := db.GetDB().QueryRow(updateUserByUserName, user.UserName, user.PhoneNumber, user.City).Scan(&updatedUser.UserName, &updatedUser.PhoneNumber, &updatedUser.City, &updatedUser.CreatedAt, &updatedUser.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}

	return updatedUser, nil
}

func AddNewUser(user model.User) (model.User, error) {

	var userCreated model.User

	err := db.GetDB().QueryRow(insertUser, user.UserName, user.PhoneNumber, user.City).Scan(&userCreated.UserName, &userCreated.PhoneNumber, &userCreated.City, &userCreated.CreatedAt, &userCreated.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}

	return userCreated, nil
}

func GetAllUsers() ([]model.User, error) {

	var users []model.User
	var username, phonenumber, city string
	var created_at, updated_at time.Time

	rows, err := db.GetDB().Query(getAllUsers)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&username, &phonenumber, &city, &created_at, &updated_at)
		if err != nil {
			return nil, err
		}
		user := model.User{
			UserName:    username,
			PhoneNumber: phonenumber,
			City:        city,
			CreatedAt:   created_at,
			UpdatedAt:   updated_at,
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByUserName(username string) (model.User, error) {

	var phonenumber, city string
	var created_at, updated_at time.Time
	row := db.GetDB().QueryRow(getUserByUserName, username)
	err := row.Scan(&phonenumber, &city, &created_at, &updated_at)
	if err != nil {
		return model.User{}, err
	}
	user := model.User{
		UserName:    username,
		PhoneNumber: phonenumber,
		City:        city,
		CreatedAt:   created_at,
		UpdatedAt:   updated_at,
	}
	return user, nil

}
