package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"unique" json:"email"`
	Name  string `json:"name"`
}

func (user *User) GetAllUsers() ([]User, error) {
	var users []User
	err := GetDB().Find(&users).Error
	return users, err
}

func (user *User) Create() error {
	err := GetDB().Create(user).Error
	return err
}

func (user *User) Delete(id uint) error {
	delUser := &User{}
	err := GetDB().Where("id=?", id).Find(delUser).Error
	if err != nil {
		log.Println(err)
	}
	return GetDB().Delete(delUser).Error
}

func (user *User) Update() error {
	return GetDB().Save(user).Error
}

func (user *User) Exists(userId uint) bool {
	err := GetDB().Where("id=?", userId).Find(&User{}).Error
	return err == nil
}
