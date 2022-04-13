package pkg

import (
	"fmt"
	"reflect"
)

func LogIn() {
	var username string
	var password string
	fmt.Println("Please enter Username: ")
	fmt.Scanln(&username)
	fmt.Println("Please enter Password: ")
	fmt.Scanln(&password)

	isAuthenticated := AuthenticateUser(username, password)

	if isAuthenticated {
		fmt.Println("User Authentication Success!")
		GetUserInput(username)
	} else {
		fmt.Println("User Authentication Failed!!!")
		fmt.Println("Please try again")
		LogIn()
	}
}

func AuthenticateUser(username, password string) bool {
	CurrentUser := User{username, password}
	Users, err := ReadLinesScanner()
	if err != nil {
		fmt.Println("Cannot read users list: ", err)
	}
	for _, use := range Users {
		// Reference: https://pkg.go.dev/reflect#DeepEqual
		if reflect.DeepEqual(use, CurrentUser) {
			return true
		}
	}
	return false
}

func SignUp() {
	var username string
	var password string
	fmt.Println("Please enter Username: ")
	fmt.Scanln(&username)
	fmt.Println("Please enter Password: ")
	fmt.Scanln(&password)

	AddNewUser(username, password)
	GetUserInput(username)
}
