package utils

import (
	"fmt"
	"log"
	"strings"
)

type User struct {
	Username string
	Password string
}

func Login() {
	var username string
	var password string
	fmt.Println("Please enter username : ")
	fmt.Scanln(&username)
	fmt.Println("Please enter password : ")
	fmt.Scanln(&password)

	isAuthenticated := Authenticate(username, password)

	if isAuthenticated {
		fmt.Println("\nNow you are logged in!!")
		UserInput(username)
	} else {
		fmt.Println("\nInvalid Username or Password!")
		Login()
	}
}

func Authenticate(username, password string) bool {

	Users, err := ReadFromFile(USERS_PATH)

	if err != nil {
		return false
	}

	for _, value := range Users {
		user := strings.Fields(value)
		if user[0] == username && user[1] == password {
			return true
		}
	}

	return false
}

func Signup() {
	var username, password string
	
	users, e := ReadFromFile(USERS_PATH)
	
	if e != nil {
		log.Fatal(e)
		return
	}
	
	if len(users) >= 2 {
		fmt.Println("Sorry, users limit reached!!")
		return
	}
	
	fmt.Println("Please enter username : ")
	fmt.Scanln(&username)
	fmt.Println("Please enter password : ")
	fmt.Scanln(&password)
	
	AddUser(username, password)
	UserInput(username)

}
