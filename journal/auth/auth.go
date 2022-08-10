package auth

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"syscall"

	"journal/secure"
	"journal/journal_handler"
	"golang.org/x/term"
)

const (
	USERS_COLLECTION = "accounts.txt"
	JOURNALS_COLLECTION = "database"
	USERS_LIMIT = 10
)

func Login(show bool, add, username, password string) {
	if username == "" && password == "" {
		username,password = getUserCreds(username,password)
	}
	if _, err := os.Stat(USERS_COLLECTION); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("User not found :(")
		return
	}
	
	var txt []byte
	txt, err := ioutil.ReadFile(USERS_COLLECTION)
	if err != nil {
		fmt.Println("Error while reading USERS_COLLECTION, err: ", err)
		return
	}
	if txt == nil || len(txt) < 1 {
		fmt.Printf("User not found :(")
		return
	}
	lines := strings.Split(string(txt), "\n")
	var exist bool
	for _, val := range lines {
		if val == "" {
			break
		}
		line := secure.Decode(val)
		user := strings.Split(line, " ")
		if user[0] == username {
			if user[1] == password {
				exist = true
				break
			}
		}
	}
	if !exist {
		fmt.Printf("User not found :(\n")
		return
	}
	fmt.Printf("\nLogged In :)\n")
	if(len(add) > 0) {
		journal_handler.AddJournal(add, username)
	} else if(show) {
		journal_handler.ListJournals(username)
	} else {
		journal_handler.TaskSelection(username)
	}
}

func Register(username, password string) {
	if username == "" && password == "" {
		username,password = getUserCreds(username,password)
	}

	text := username + " " + password
	accountLimit := journal_handler.VerifyLimit(USERS_COLLECTION, USERS_LIMIT)
	if !accountLimit {
		return
	}
	file, err := os.OpenFile(USERS_COLLECTION, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Error while reading USERS_COLLECTION, err: ", err)
		return
	}
	defer file.Close()

	encryptedEntry := secure.Encode(text) + "\n"
	if _, err = file.WriteString(encryptedEntry); err != nil {
		fmt.Println("Error while writing to USERS_COLLECTION, err: ", err)
		return
	}
	fmt.Printf("\nSigned In :)\n")

	homeDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	journalDir := homeDir + "/" + JOURNALS_COLLECTION + "/"
	fileName := journalDir + username + ".txt"

	if _, err := os.Stat(journalDir); os.IsNotExist(err) {
		err := os.Mkdir(journalDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		_, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
	}
	journal_handler.TaskSelection(username)
}

func getUserCreds(username, password string) (string,string) {
	fmt.Print("Enter Username:")
	fmt.Scanln(&username)
	fmt.Print("Enter Password:")
	bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
	password = string(bytePassword)
	return username,password
}