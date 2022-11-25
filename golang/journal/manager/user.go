package manager

import (
	"fmt"
	"journal/helper"
	"journa;/crypt"
	"os"
)

func CaptureUser(name string, email string, password string) error {
	fmt.Println(name, email, password)
	err := helper.CheckFile("file")
	if err != nil {
		fmt.Println(err)
		return err
	}

	bytes, err := os.ReadFile("file")
	if err != nil {
		return err
	}
	decrypt,err := crypt.Decrypt(string(bytes))

	return nil
}

func LogIn(email string, password string) {
	fmt.Println(email, password)
}

func LogOut(email string) {
	fmt.Println(email)
}

// func LoadUsers(){

// }
