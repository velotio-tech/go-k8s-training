package users

import (
	"fmt"
	"io/ioutil"
	"journal/constants"
	"journal/encrypt"
	"log"
	"os"
	"strings"
	"time"
)

type User struct {
	Username string
	Password string
}

func (u *User) ReadJournal() []string {
	file := constants.DB_LOCATION + u.Username + constants.JOURNAL_NAME
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
			fmt.Println(cipherText[0], string(plainText))
			text = append(text, cipherText[0])
			text = append(text, string(plainText))
		}
		return text
	}
}

func (u *User) WriteJournal(entry string) {
	file := constants.DB_LOCATION + u.Username + constants.JOURNAL_NAME
	data := u.ReadJournal()
	var journal = make([]string, constants.MAX_JOURNAL_ENTRY, constants.MAX_JOURNAL_ENTRY)
	journal = append(journal, data...)
	journal = append(journal, entry)
	if len(journal) > constants.MAX_JOURNAL_ENTRY {
		diff := len(journal) - constants.MAX_JOURNAL_ENTRY
		journal = journal[diff:]
	}
	var str string
	for _, wj := range journal {
		str += wj + "\n"
	}
	f, _ := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	cipherText, _ := encrypt.Encrypt([]byte(entry))
	time := time.Now()
	currentTime := time.Format("02 Jan 06 15:04 MST -")
	cipherText = append([]byte(currentTime), cipherText...)
	cipherText = append(cipherText, "\n"...)
	_, err := f.Write(cipherText)
	if err != nil {
		fmt.Println("Failed to add entry to journal")
		return
	}
	defer f.Close()
}

func create(usr, pswd string) (*User, error) {
	var total, _ = getUserCount()
	var err error
	if total < constants.MAX_USERS {
		fmt.Printf("Setting up user : %s \n", usr)
		addUser(usr, pswd)
		fmt.Printf("Creating new journal for %s \n", usr)
		createJournal(usr + constants.JOURNAL_NAME)
		u := User{
			Username: usr,
			Password: pswd,
		}
		return &u, nil
	} else {
		fmt.Printf("User limit exceeded. Increase limit to create new user \n")
		return nil, err
	}
}

func createJournal(s string) {
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

func exists(u string) bool {
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

func authenticate(u, p string) bool {
	uMap := getUsers()
	isExists := exists(u)
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

func AddEntry(usr, pswd, entry string) {
	if authenticate(usr, pswd) {
		fmt.Println("User Authenticated.....")
		entryUser := &User{
			Username: usr,
			Password: pswd,
		}
		if entryUser != nil {
			entryUser.WriteJournal(entry)
		} else {
			fmt.Println("Cannot add entry")
		}
	} else {
		fmt.Println("Incorrect password. Please try again.....")
	}
}

func ShowEntry(usr, pswd string) {
	if authenticate(usr, pswd) {
		fmt.Println("User Authenticated.....")
		entryUser := &User{
			Username: usr,
			Password: pswd,
		}
		if entryUser != nil {
			entryUser.ReadJournal()
		}
	} else {
		fmt.Println("Incorrect password. Please try again")
	}
}

func LoginUser(username, password string) *User {
	if authenticate(username, password) {
		authUser := &User{
			Username: username,
			Password: password,
		}
		return authUser
	} else {
		return nil
	}
}

func Signup(username, password string) *User {
	if exists(username) {
		fmt.Println("User already exists. Login or Use different username")
		return new(User)
	} else {
		newUser, err := create(username, password)
		if err == nil {
			fmt.Printf("User account for username : %s created! \n", username)
			return newUser
		} else {
			log.Fatalf("Error : %s", err)
			return newUser
		}
	}
}
