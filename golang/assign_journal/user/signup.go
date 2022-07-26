package user

import (
	"fmt"
	"os"
	"syscall"

	"github.com/6akshita/assign_journal/utils"
	"golang.org/x/term"
)

func UserSignUp(username, password string) {
	if username == "" && password == "" {
		fmt.Print("Enter Username:")
		fmt.Scanln(&username)
		fmt.Print("Enter Password:")
		bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
		password = string(bytePassword)
	}

	text := username + " " + password
	res := utils.UserLimitCheck("users.txt", 10)
	if !res {
		return
	}
	file, err := os.OpenFile("users.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Error while opening Users file, err: ", err)
		return
	}
	defer file.Close()

	encryptedEntry := utils.Encrypt(text) + "\n"
	if _, err = file.WriteString(encryptedEntry); err != nil {
		fmt.Println("Error while adding user entry to file, err: ", err)
		return
	}
	fmt.Printf("User successfully signed up to the app!!\n")

	homeDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	journalDir := homeDir + "/journal/"
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
	utils.EntryCheck(username)
}
