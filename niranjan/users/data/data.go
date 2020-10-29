package models

import (
	"encoding/json"
	"io"
)

//User is a struct containing user fields in db
type User struct {
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first name"`
	LastName  string `json:"last name"`
	Phone     int    `json:"phone"`
}

//Users is a list of User records
type Users []*User

//ToJson converts to json
func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

// GetUsers returns a list of users
func (u *Users) GetUsers() Users {

	//Fetch user list from DB and store to userList

	return userList
}
