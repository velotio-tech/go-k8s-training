package cmd

import "time"

type User struct {
	Name     string
	Email    string
	Password string
	Journal  []*Entry
}

func NewUser(name, email, password string) *User {
	return &User{Name: name, Email: email, Password: password}
}

type Entry struct {
	Id        int
	CreatedAt string
	Message   string
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
