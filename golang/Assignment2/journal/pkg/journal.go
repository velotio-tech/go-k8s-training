package pkg

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	Password    string   `json:"password"`
	LoginStatus bool     `json:"loginStatus"`
	Journal     []*Entry `json:"journal"`
}

func NewUser(name, email, password string) *User {
	return &User{Name: name, Email: email, Password: password, Journal: make([]*Entry, 0)}
}

func (u *User) AddEntry(entry *Entry) {
	u.Journal = append(u.Journal, entry)
	// Evict the oldest entry
	if len(u.Journal) > 50 {
		fmt.Println("Journal size greater than 50, evicting the oldest entry")
		u.Journal = u.Journal[1:len(u.Journal)]
	}
}

type Entry struct {
	Id        int    `json:"id"`
	CreatedAt string `json:"createdAt"`
	Message   string `json:"message"`
}

func NewEntry(message string) *Entry {
	return &Entry{Id: entryId(), CreatedAt: time.Now().Format("02 Jan 2006 03:04PM"), Message: message}
}

func (e *Entry) IsValid() bool {
	return strings.TrimSpace(e.Message) != ""
}

var entryId = IdGenerator()

func IdGenerator() func() int {
	id := 0
	return func() int {
		id += 1
		return id
	}
}
