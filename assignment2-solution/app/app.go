package app

import "fmt"

type EntryData struct {
	time []string
	data []string
}

type User struct {
	username string
	password string
	email    string
}

var entries []string

var userMap = make(map[string]EntryData)

func ValidateUser(uname string) bool {
	if _, present := userMap[uname]; present {
		return true
	} else {
		fmt.Println("Unauthorized user, Please register!")
		return false
	}
}

func AddNewUserToMap(username string) {
	var data EntryData

	if len(userMap) > 10 {
		fmt.Println("Entries exceeded!")
		return
	}

	userMap[username] = data
	fmt.Println("New user added successfully!!")
}
