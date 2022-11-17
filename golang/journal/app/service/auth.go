package service

import (
	"encoding/json"
	"fmt"

	"github.com/segmentio/ksuid"
	"github.com/velotio-ajaykumbhar/journal/app/model"
	"github.com/velotio-ajaykumbhar/journal/util"
)

var (
	AuthFilename    = "auth.lock"
	SessionFilename = "session.lock"
)

func accountExists(username string, auths []model.Auth) (model.Auth, bool) {
	for _, auth := range auths {
		if auth.Username == username {
			return auth, true
		}
	}
	return model.Auth{}, false
}

func authValidation(username, passowrd string) error {
	if username == "" && passowrd == "" {
		return fmt.Errorf("username and password are required field")
	}

	if username == "" {
		return fmt.Errorf("username is required field")
	}

	if passowrd == "" {
		return fmt.Errorf("password is required field")
	}

	return nil
}

func SignUp(username, password string) {

	if err := authValidation(username, password); err != nil {
		fmt.Println(err.Error())
		return
	}

	data, err := util.ReadFile(AuthFilename)
	if err != nil {
		fmt.Println("error while sign up. please try again")
		return
	}

	data = util.Decode(string(data))

	auths := []model.Auth{}
	json.Unmarshal(data, &auths)

	_, exists := accountExists(username, auths)
	if exists {
		fmt.Println("account already exists")
		return
	}

	auth := model.Auth{
		UserId:   ksuid.New().String(),
		Username: username,
		Password: password,
	}

	auths = append(auths, auth)
	marshal, _ := json.Marshal(auths)

	encode := util.Encode(marshal)

	err = util.WriteFile(AuthFilename, []byte(encode))
	if err != nil {
		fmt.Println("error while sign up. please try again")
	}

	setSession(auth)
	fmt.Println("account created successfully")
}

func Login(username, password string) {

	if err := authValidation(username, password); err != nil {
		fmt.Println(err.Error())
		return
	}

	data, err := util.ReadFile(AuthFilename)
	if err != nil {
		fmt.Println("error while sign up. please try again")
		return
	}
	data = util.Decode(string(data))

	auths := []model.Auth{}
	json.Unmarshal(data, &auths)

	auth, exists := accountExists(username, auths)
	if !exists {
		fmt.Println("account not exists")
		return
	}

	setSession(auth)
	fmt.Println("login successfull")
}

func Logout() {
	auth := model.Auth{}
	setSession(auth)
}
