package database

import(
	"encoding/json"
	"fmt"
	"os"
)

const DB_FILENAME = "database/db.json"

type User struct {
	Username string
	Name string
}

func getUsers() ([]User, error) {
	data, err := os.ReadFile(DB_FILENAME)
	var users []User

	if err == nil {
		json.Unmarshal(data, &users)
	}
	
	return users, err
}

func FindUser(username string) (*User, error) {
	users, err := getUsers()

	if(err == nil) {
		for index := 0; index < len(users); index++ {
			user := users[index]

			if(user.Username == username) {
				return &user, nil
			}
		}
	}

	return nil, err
}

func updateDB(data []User) {
	bytes, err := json.Marshal(data)
		
	if err == nil {
		os.WriteFile(DB_FILENAME, bytes, 0644)
	} else {
		fmt.Println("Failed to save to DB file '%s'", err)
	}
}

func CreateUser(username string, name string) (*User, error) {
	user, err := FindUser(username)

	if user != nil {
		return user, err
	}
	
	newUser := User{
		Username: username,
		Name: name,
	}

	users, err := getUsers()
	
	if err == nil {
		users = append(users, newUser)
		updateDB(users)
	}

	return &newUser, err
}

