package cmd

import "time"

type User struct {
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	Password    string   `json:"password"`
	LoginStatus bool     `json:"loginStatus"`
	Journal     []*Entry `json:"journal"`
}

func NewUser(name, email, password string) *User {
	return &User{Name: name, Email: email, Password: password}
}

type Entry struct {
	Id        int    `json:"id"`
	CreatedAt string `json:"createdAt"`
	Message   string `json:"message"`
}

func NewEntry(message string) *Entry {
	return &Entry{Id: entryId(), CreatedAt: time.Now().Format("02 Jan 2006 03:04PM"), Message: message}
}

var entryId = IdGenerator()

func IdGenerator() func() int {
	id := 0
	return func() int {
		id += 1
		return id
	}
}
