package users

import (
	"fmt"
	"io/ioutil"
	"journal/constants"
	"journal/encrypt"
	"os"
	"strings"
	"time"
)

type User struct {
	username string
	password string
}

func Create(u, p string) (*User, error) {
	var total, _ = getUserCount()
	var err error
	if total < constants.MAX_USERS {
		fmt.Printf("Setting up user : %s \n", u)
		addUser(u, p)
		fmt.Printf("Creating new journal for %s \n", u)
		CreateJournal(u + constants.JOURNAL_NAME)
		u := User{
			username: u,
			password: p,
		}
		return &u, nil
	} else {
		fmt.Printf("User limit exceeded. Increase limit to create new user \n")
		return nil, err
	}
}

func CreateJournal(s string) {
	journal, err := os.Create(constants.DB_LOCATION + s)
	if err != nil {
		fmt.Println("Failed to create journal", err)
	} else {
		fmt.Println("Successfully created new journal. You can add entry with : journal entry --add")
	}
	defer journal.Close()
}

func getUserCount() (int, []byte) {
	text, err := ioutil.ReadFile(constants.DB_LOCATION + constants.USER_DB)
	if err != nil {
		fmt.Printf("%s not found at %s", constants.USER_DB, constants.DB_LOCATION)
		return len(text), nil
	} else {
		total := strings.Split(string(text), "\n")
		return len(total), text
	}
}

func addUser(u, p string) {
	fileName := constants.DB_LOCATION + constants.USER_DB
	file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	entry := u + "," + p
	_, err := file.Write([]byte(entry + "\n"))
	if err != nil {
		fmt.Println("Failed to add user \n", err)
		return
	} else {
		fmt.Printf("Successfully added new user : %s. \n", u)
	}
	defer file.Close()
}

func GetValue(username, password string) *User {
	var loginUser = User{
		username: username,
		password: password,
	}
	return &loginUser
}

func Exists(u string) bool {
	uMap := getUsers()
	if _, ok := uMap[u]; ok {
		return true
	} else {
		return false
	}
}

func getUsers() map[string]string {
	count, users := getUserCount()
	var uToP = map[string]string{}
	if users == nil || count == 0 {
		return uToP
	}
	splitUser := strings.Split(string(users), "\n")
	for _, s := range splitUser {
		if s == "" {
			break
		}
		uMap := strings.Split(s, ",")
		uToP[uMap[0]] = uMap[1]
	}
	return uToP
}

func Auth(u, p string) bool {
	uMap := getUsers()
	isExists := Exists(u)
	if isExists {
		value := uMap[u]
		if p == value {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (u *User) JournalWrite(entry string) {
	file := constants.DB_LOCATION + u.username + constants.JOURNAL_NAME
	data := journalRead(file)
	var w = make([]string, 50, 50)
	w = append(w, data...)
	w = append(w, entry)
	if len(w) > constants.MAX_JOURNAL_ENTRY {
		diff := len(w) - constants.MAX_JOURNAL_ENTRY
		w = w[diff:]
	}
	var s string
	for _, wj := range w {
		s += wj + "\n"
	}
	f, _ := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	cipherText, _ := encrypt.Encrypt([]byte(entry))
	n := "\n"
	time := time.Now()
	currentTime := time.Format("02 Jan 06 15:04 MST -")
	cipherText = append([]byte(currentTime), cipherText...)
	cipherText = append(cipherText, n...)
	_, err := f.Write(cipherText)
	if err != nil {
		fmt.Println("Failed to add entry to journal")
		return
	}
	defer f.Close()
}

func journalRead(file string) []string {
	data, err := ioutil.ReadFile(file)
	var text []string
	if err != nil {
		fmt.Println("File not found")
		return make([]string, 0)
	} else {
		if len(data) < 1 {
			return text
		}
		lines := strings.Split(string(data), "\n")
		for _, l := range lines {
			if l == "" {
				break
			}
			cipherText := strings.Split(l, "-")
			plainText, _ := encrypt.Decrypt([]byte(cipherText[1]))
			text = append(text, string(cipherText[0]))
			text = append(text, string(plainText))
		}
		fmt.Print(text)
		return text
	}
}
