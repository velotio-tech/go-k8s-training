package manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"journal/crypt"
	"journal/helper"
	"os"
	"strings"
)

func CaptureNewUser(name string, email string, password string) error {
	err := helper.CheckFile("registeredusers")
	helper.Check(err)

	cipher, err := os.ReadFile("registeredusers")
	helper.Check(err)

	data, err := helper.ByteToSliceOfStrings(cipher)
	helper.Check(err)

	isDuplicate := helper.CheckIfDuplicateUser(data, name)
	if isDuplicate {
		return errors.New("User already exist")
	}

	data = append(data, name)
	fmt.Println(data)
	byteData := helper.SliceOfStringsToByte(data)

	helper.EncryptAndAppendToFile(byteData, "registeredusers")

	CreateUserEntry(name, email, password)
	return nil
}

func CreateUserEntry(name string, email string, password string) {
	err := helper.CheckFile("data")
	helper.Check(err)

	cipher, err := os.ReadFile("data")
	helper.Check(err)

	var newStr string
	user := NewUser(name, email, password)
	d, _ := json.Marshal(user)

	if len(cipher) != 0 {
		str, _ := crypt.Decrypt(string(cipher))
		newStr = str[:len(str)-1] + "," + string(d) + "]"
	} else {
		newStr = "[" + string(d) + "]"
	}
	fmt.Println(newStr)

	helper.EncryptAndAppendToFile([]byte(newStr), "data")
}

func LogIn(email string, password string) error {
	// email should already exist in "data" file
	err := helper.CheckFile("data")
	helper.Check(err)

	cipher, err := os.ReadFile("data")
	helper.Check(err)

	u := User{}
	var newData string = ""
	var count int = 0
	if len(cipher) != 0 {
		str, _ := crypt.Decrypt(string(cipher))
		strSlice := strings.Split(str[1:len(str)-1], "},{")
		for index, data := range strSlice {
			if index == 0 {
				data = data + "}"
			} else if index == len(strSlice)-1 {
				data = "{" + data
			} else {
				data = "{" + data + "}"
			}
			json.Unmarshal([]byte(data), &u)
			if u.Email == email && u.Password == password {
				u.IsLoggedIn = true
				s, _ := json.Marshal(u)
				data = string(s)
				count = count + 1
			}
			newData = newData + "," + data
		}
		if count == 0 {
			return errors.New("Email does not exits")
		}
		newData = "[" + newData[1:] + "]"
		helper.EncryptAndAppendToFile([]byte(newData), "data")
		fmt.Println(newData)
	} else {
		return errors.New("Email does not exits")
	}
	return nil
}
