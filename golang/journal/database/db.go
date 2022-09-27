package database

import(
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const USER_DB_FILENAME = "database/users.json"
const ENTRY_DB_FILENAME = "database/entries.json"

type User struct {
	Username string
	Name string
}

type Entry struct {
	Username string
	Content string
	InsertedAt time.Time
}

func getUsers() ([]User, error) {
	data, err := os.ReadFile(USER_DB_FILENAME)
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

func updateUsersDB(data []User) {
	bytes, err := json.Marshal(data)
		
	if err == nil {
		os.WriteFile(USER_DB_FILENAME, bytes, 0644)
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
		updateUsersDB(users)
	}

	return &newUser, err
}

func getEntries() ([]Entry, error) {
	data, err := os.ReadFile(ENTRY_DB_FILENAME)

	var entries []Entry

	if err == nil {
		json.Unmarshal(data, &entries)
	}

	return entries, err
}

func updateEntriesDB(data []Entry) error {
	bytes, err := json.Marshal(data)

	if err == nil {
		os.WriteFile(ENTRY_DB_FILENAME, bytes, 0644)
	} else {
		fmt.Println("Failed to save to DB file '%s'", err)
	}

	return err
}

func CreateEntry(username string, content string) (*Entry, error) {
	newEntry := Entry{
		Username: username,
		Content: content,
		InsertedAt: time.Now(),
	}

	entries, err := getEntries()

	if err == nil {
		entries := append(entries, newEntry)
		err = updateEntriesDB(entries)
	}

	return &newEntry, err
}

