pacakge database

import(
	"encoding/json"
	"os"
	"strings"
)

type User struct {
	username string
	name string
	password string
}

func CreateUser(username string, name string, password string) *User {
	newUser := User{
		username: username,
		name: name,
		password: password,	
	}

	return newUser
}

