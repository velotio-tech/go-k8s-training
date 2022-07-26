package user

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"syscall"

	"github.com/6akshita/assign_journal/utils"
	"golang.org/x/term"
)

func LoginUser(username, password string) {
	if username == "" && password == "" {
		fmt.Print("Enter Username:")
		fmt.Scanln(&username)
		fmt.Print("Enter Password:")
		bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
		password = string(bytePassword)
	}
	if _, err := os.Stat("users.txt"); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("User doesn't exist, Sign up first!")
		return
	}
	// Get all users
	var txt []byte
	txt, err := ioutil.ReadFile("users.txt")
	if err != nil {
		fmt.Println("Error while reading users file, err: ", err)
		return
	}
	if txt == nil || len(txt) < 1 {
		fmt.Printf("User doesn't exist, Sign up first!")
		return
	}
	lines := strings.Split(string(txt), "\n")
	var userFound bool
	for _, val := range lines {
		if val == "" {
			break
		}
		line := utils.Decrypt(val)
		user := strings.Split(line, " ")
		if user[0] == username {
			if user[1] == password {
				userFound = true
				break
			}
		}
	}
	if !userFound {
		fmt.Printf("User doesn't exist, Sign up first!\n")
		return
	}
	fmt.Printf("User logged in successfully!\n")
	utils.EntryCheck(username)
}
