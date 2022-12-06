package manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"journal/crypt"
	"journal/helper"
	"os"
)

type Journal struct {
	CreatedAt string `json:"createdAt"`
	Message   string `json:"message"`
}

type User struct {
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	IsLoggedIn bool      `json:"isloggedin"`
	Journal    []Journal `json:"journal"`
}

type users []User

func NewUser(name, email, password string) User {
	// &variable = address
	// *address = variable
	u := User{Name: name, Email: email, Password: password, IsLoggedIn: false, Journal: make([]Journal, 0)}
	return u
}

func AddEntry(message string, email string) error {
	err := helper.CheckFile("data")
	helper.Check(err)

	cipher, err := os.ReadFile("data")
	helper.Check(err)

	u := users{}
	j := Journal{CreatedAt: "today", Message: message}
	var count int = 0
	if len(cipher) != 0 {
		str, _ := crypt.Decrypt(string(cipher))
		json.Unmarshal([]byte(str), &u)
		for index, data := range u {
			if data.Email == email && data.IsLoggedIn {
				count = count + 1
				u[index].Journal = append(u[index].Journal, j)
				fmt.Println(u[index])
			}
		}
		newData, _ := json.Marshal(u)

		helper.EncryptAndAppendToFile(newData, "data")
		if count == 0 {
			return errors.New("Email does not exist or you are not logged in")
		}
	} else {
		return errors.New("Email does not exist")
	}

	return nil
}

func ListEntry(email string) error {
	err := helper.CheckFile("data")
	helper.Check(err)

	cipher, err := os.ReadFile("data")
	helper.Check(err)

	u := users{}
	if len(cipher) != 0 {
		str, _ := crypt.Decrypt(string(cipher))
		json.Unmarshal([]byte(str), &u)
		for _, data := range u {
			if data.Email == email && data.IsLoggedIn {
				for _, journal := range data.Journal {
					j, _ := json.Marshal(journal)
					fmt.Println(string(j))
				}
			}
		}
	}
	return nil
}
