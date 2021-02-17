package models

import (
	"regexp"

	"github.com/pankaj9310/go-k8s-training/pankaj/golang/assignment3/user/api/utils"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) Validate() (map[string]interface{}, bool) {

	if len(user.Email) == 0 {
		return utils.Message(false, "Email address is required"), false
	} else {
		var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !emailRegex.MatchString(user.Email) {
			return utils.Message(false, "Email address is not Valid"), false
		}
	}

	if len(user.Password) < 6 {
		return utils.Message(false, "Password is required"), false
	}

	//Email must be unique
	temp := &User{}

	//check for errors and duplicate emails
	err := GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return utils.Message(false, "Email address already in use by another user."), false
	}

	return utils.Message(false, "Requirement passed"), true
}

func (user *User) Create() map[string]interface{} {

	if resp, ok := user.Validate(); !ok {
		return resp
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return utils.Message(false, "Invalid password")
	}
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	if user.ID <= 0 {
		return utils.Message(false, "Failed to create user, connection error.")
	}

	user.Password = "" //delete password

	response := utils.Message(true, "Successfully Created User")
	response["user"] = user
	return response
}
func (user *User) Update(id string) map[string]interface{} {
	temp := &User{}

	err := GetDB().Where("id=?", id).First(temp).Error
	if err != nil {

		return utils.Message(true, "User Not Found!")
	}

	if user.Password != "" {
		if len(user.Password) < 6 {
			return utils.Message(false, "Password is required")
		} else {
			hashedPassword, err := utils.HashPassword(user.Password)
			if err != nil {
				return utils.Message(false, "Invalid password")
			}
			temp.Password = string(hashedPassword)

		}
	}
	temp.Name = user.Name
	GetDB().Save(temp)
	user.Password = "" //delete password

	response := utils.Message(true, "Successfully Updated User")
	response["user"] = user
	return response
}

func GetUser(id string) *User {

	user := &User{}
	GetDB().Where("id = ?", id).First(user)
	if user.Email == "" { //User not found!
		return nil
	}

	user.Password = ""
	return user
}

func DeleteUser(id string) map[string]interface{} {
	user := &User{}
	err := GetDB().Where("id=?", id).Find(user).Error
	if err != nil {
		return utils.Message(false, "User Not Found!")
	}
	GetDB().Delete(user)
	return utils.Message(true, "Successfully Deleted User")
}
