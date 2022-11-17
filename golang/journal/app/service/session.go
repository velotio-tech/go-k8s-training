package service

import (
	"encoding/json"
	"fmt"

	"github.com/velotio-ajaykumbhar/journal/app/model"
	"github.com/velotio-ajaykumbhar/journal/util"
)

func setSession(auth model.Auth) {
	marshal, _ := json.Marshal(auth)

	encode := util.Encode(marshal)

	err := util.WriteFile(SessionFilename, []byte(encode))
	if err != nil {
		fmt.Println("error while sign up. please try again")
	}
}

func GetSession() model.Auth {
	data, err := util.ReadFile(SessionFilename)
	if err != nil {
		fmt.Println("error while sign up. please try again")
		return model.Auth{}
	}
	data = util.Decode(string(data))
	auth := model.Auth{}
	json.Unmarshal(data, &auth)
	return auth
}
